package main

import (
	"log"

	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/router"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm/logger"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//"be_latihan/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
    AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
    AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
    AllowHeaders: "Origin,Content-Type,Accept,Authorization",
}))

	config.InitDB()
	config.GetDB().AutoMigrate(&model.Mahasiswa{}, &model.User{})
	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}