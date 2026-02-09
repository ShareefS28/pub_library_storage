package migrations

import "gorm.io/gorm"

func CreateSessionCleanUpJob(db *gorm.DB) error {
	sql := `
		USE msdb;

		IF NOT EXISTS (
			SELECT 1
			FROM msdb.dbo.sysjobs
			WHERE name = N'CleanupExpiredSessions'
		)
		BEGIN
        EXEC sp_add_job
                @job_name = N'CleanupExpiredSessions',
                @enabled = 1,
                @description = N'Delete expired or revoked auth sessions',
                @owner_login_name = N'sa';

        EXEC sp_add_jobstep
                @job_name = N'CleanupExpiredSessions',
                @step_name = N'Delete expired sessions',
                @subsystem = N'TSQL',
                @database_name = N'library',
                @command = N'
                        DELETE FROM library.dbo.mt_sessions
                        WHERE expired_at < SYSUTCDATETIME()
                           OR revoked_at IS NOT NULL;
                ',
                @retry_attempts = 3,
                @retry_interval = 5;

        EXEC sp_add_schedule
                @schedule_name = N'CleanupExpiredSessions_Every30Min',
				@enabled = 1,
                @freq_type = 4,
                @freq_interval = 1,
                @freq_subday_type = 4,
                @freq_subday_interval = 30; -- every 30 min

        EXEC sp_attach_schedule
			@job_name = N'CleanupExpiredSessions',
			@schedule_name = N'CleanupExpiredSessions_Every30Min';

        EXEC sp_add_jobserver
            @job_name = N'CleanupExpiredSessions';
	END
	`

	return db.Exec(sql).Error
}

func DropSessionCleanupJob(db *gorm.DB) error {
	sql := `
		USE msdb;

		IF EXISTS (
			SELECT 1
			FROM msdb.dbo.sysjobs
			WHERE name = N'CleanupExpiredSessions'
		)
		BEGIN
			EXEC sp_update_job
				@job_name = N'CleanupExpiredSessions',
				@enabled = 0;
			
			EXEC sp_delete_job
				@job_name = N'CleanupExpiredSessions',
				@delete_unused_schedule = 1;
		END
	`

	return db.Exec(sql).Error
}

/*
SELECT
    job_id,
    name AS job_name,
    enabled,
    date_created,
    date_modified
FROM msdb.dbo.sysjobs
ORDER BY name;

SELECT
    j.name AS job_name,
    h.run_date,
    h.run_time,
    h.run_duration,
    h.run_status,
    CASE h.run_status
        WHEN 0 THEN 'Failed'
        WHEN 1 THEN 'Succeeded'
        WHEN 2 THEN 'Retry'
        WHEN 3 THEN 'Canceled'
        WHEN 4 THEN 'In Progress'
    END AS status
FROM msdb.dbo.sysjobs j
JOIN msdb.dbo.sysjobhistory h
    ON j.job_id = h.job_id
WHERE j.name = 'CleanupExpiredSessions'
  AND h.step_id = 0
ORDER BY h.run_date DESC, h.run_time DESC;

EXEC msdb.dbo.sp_help_jobhistory
@job_name = N'CleanupExpiredSessions';

SELECT COUNT(*) AS rows_to_delete
FROM library.dbo.mt_sessions
WHERE expired_at < SYSUTCDATETIME()
   OR revoked_at IS NOT NULL;

EXEC msdb.dbo.sp_start_job
	@job_name = N'CleanupExpiredSessions';

SELECT servicename, status_desc
FROM sys.dm_server_services
WHERE servicename LIKE '%Agent%';

SELECT SERVERPROPERTY('Edition') AS edition;

*/
