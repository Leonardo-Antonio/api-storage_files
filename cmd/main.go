package main

import (
	"github.com/Leonardo-Antonio/api-storage_files/pkg/db"
	"github.com/Leonardo-Antonio/api-storage_files/pkg/image"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rds := db.RedisGetClient()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	image.Router(app, rds)
	app.Listen(":8000")
}
