package router

import (
	"inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") 

	// Route untuk homepage
	api.Get("/", handler.Homepage)

	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM, middleware.Middlewares("admin"))

	// Route untuk menambah mahasiswa baru
	api.Post("/mahasiswa", handler.InsertMahasiswa, middleware.Middlewares("admin"))

	// Route untuk mengupdate data mahasiswa berdasarkan NPM
	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa, middleware.Middlewares("admin"))

	// Route untuk menghapus data mahasiswa berdasarkan NPM
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa, middleware.Middlewares("admin"))

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// swagger documentation
	app.Get("/docs/*", swagger.HandlerDefault)

}