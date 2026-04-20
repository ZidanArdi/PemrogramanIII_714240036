package router
import (
	"be_latihan/model"
	"github.com/gofiber/fiber/v2"
	"be_latihan/handler"
)	
func SetupRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Get("/", handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", handler.GetMahasiswaByNPM)

	//jika ingin menggunakan query parameter
	mahasiswa.Get("/search", handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", handler.InsertMahasiswa)
	mahasiswa.Put("/:npm", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", handler.DeleteMahasiswa)
}