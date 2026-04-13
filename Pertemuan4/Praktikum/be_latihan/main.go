package main

import (
	"log"
	"net/http"
	"strconv"

	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/repository"
	"be_latihan/router"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm/logger"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//"be_latihan/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	config.InitDB()
	config.GetDB().AutoMigrate(&model.Mahasiswa{})
	router .SetupRouter(app)

	app.Listen(":3000")




	app.Get("/mahasiswa", func(c *fiber.Ctx) error {
		data, err := repository.GetAllMahasiswa()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(data)
	})

	app.Get("/mahasiswa/:npm", func(c *fiber.Ctx) error {
		npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "npm tidak valid"})
		}

		data, err := repository.GetMahasiswaByNPM(npm)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(data)
	})

	app.Post("/mahasiswa", func(c *fiber.Ctx) error {
		var mhs model.Mahasiswa
		if err := c.BodyParser(&mhs); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		data, err := repository.InsertMahasiswa(&mhs)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(http.StatusCreated).JSON(data)
	})

	app.Put("/mahasiswa/:npm", func(c *fiber.Ctx) error {
		npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "npm tidak valid"})
		}

		var mhs model.Mahasiswa
		if err := c.BodyParser(&mhs); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		data, err := repository.UpdateMahasiswa(npm, &mhs)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(data)
	})

	app.Delete("/mahasiswa/:npm", func(c *fiber.Ctx) error {
		npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "npm tidak valid"})
		}

		if err := repository.DeleteMahasiswa(npm); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(http.StatusNoContent)
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}