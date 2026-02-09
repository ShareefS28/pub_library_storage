package main

import (
	"server/cmd/server/docs"
	"server/config"
	"server/database"
	"server/internal/controllers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/swagger/v2"
)

func main() {
	config.AppConfig = config.Load()

	database.Connect(config.AppConfig.DBDSN)

	// utiljwt.SetJwtKeys(config.AppConfig.JWTPrivateKey, config.AppConfig.JWTPublicKey)

	app := fiber.New()

	// Swagger endpoint
	if !config.AppConfig.IsProd {
		docs.SwaggerInfo.Title = "backend API swagger"
		docs.SwaggerInfo.Description = "server api"
		docs.SwaggerInfo.BasePath = "/api"

		app.Get("/swg/*", swagger.HandlerDefault)
	}

	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip} ${status} ${method} ${path} (${latency})\n",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: !config.AppConfig.IsProd,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	controllers.Setup(app)

	app.Listen(":" + config.AppConfig.AppPort)
}
