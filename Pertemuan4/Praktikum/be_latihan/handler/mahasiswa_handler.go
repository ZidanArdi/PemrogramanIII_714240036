package handler

import (
	"be_latihan/model"
	"be_latihan/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "Gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Message: "Berhasil mengambil data mahasiswa",
		Data:    data,
	})
}           

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npmQuery := c.Query("npm")
	if npmQuery == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "NPM harus disertakan sebagai query parameter",
		})
	}
	npm, err := strconv.ParseInt(npmQuery, 10, 64)

	//npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "NPM tidak valid",
			Error:   err.Error(),
		}) 
	}

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "Data mahasiswa dengan NPM tersebut tidak ditemukan",
			})	
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "Gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "Berhasil mengambil data mahasiswa",
		Data:    mhs,
	})
}