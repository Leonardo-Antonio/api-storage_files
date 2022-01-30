package image

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, rds *redis.Client) {
	controller := &handler{rds}

	group := app.Group(fmt.Sprintf("%s/images", os.Getenv("VERSION_API")))
	group.Post("", controller.Save)
	group.Get("", controller.GetAll)
	group.Delete("/:name", controller.Remove)
	group.Static("/public", "static/images")
}
