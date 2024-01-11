package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/srmbackisdeveloper/booksmanagement/common"
	"os"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	// init env
	err := common.LoadEnv()
	if err != nil {
		return err
	}

	// init db
	err = common.InitDb()
	if err != nil {
		return err
	}
	// defer closing db
	defer common.CloseDb()

	// create app
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// add routes

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	app.Listen(":" + port)

	return nil
}
