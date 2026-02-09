`openssl ecparam -name prime256v1 -genkey -noout -out keys/jwt_ec_private.pem`<br>
`openssl ec -in jwt_ec_private.pem -pubout -out keys/jwt_ec_public.pem`<br>

`docker pull mcr.microsoft.com/mssql/server:2022-latest`<br>
`docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=P@ssw0rd" -p 1433:1433  --name sqlpreview -d mcr.microsoft.com/mssql/server:2022-latest`<br>

`go mod tidy`<br>
`go run ./cmd/migrate`<br>
`go run ./cmd/server`<br>