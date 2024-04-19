package router

import (
	"github.com/NikSchaefer/go-fiber/handlers"
	"github.com/NikSchaefer/go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {

	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)
	apiV1Inventario := router.Group("/api/v1/inventario")
	//categoria
	categoria := apiV1Inventario.Group("/categorias")
	categoria.Post("/", handlers.CreateCategory)
	categoria.Delete("/:id", handlers.DeleteCategory)
	categoria.Get("/:id", handlers.GetCategoryById)
	categoria.Get("/", handlers.GetAllCategories)
	categoria.Put("/:id", handlers.UpdateCategory)
	//marca
	marca := apiV1Inventario.Group("/marcas")
	marca.Post("/", handlers.CreateMarca)
	marca.Delete("/:id", handlers.DeleteMarca)
	marca.Get("/:id", handlers.GetMarcaById)
	marca.Get("/", handlers.GetAllMarcas)
	marca.Put("/:id", handlers.UpdateMarca)
	//productos
	producto := apiV1Inventario.Group("/productos")
	producto.Post("/", handlers.CreateProduct)
	producto.Delete("/:id", handlers.DeleteProduct)
	producto.Get("/:id", handlers.GetProductById)
	producto.Get("/", handlers.GetAllProducts)
	producto.Put("/:id", handlers.UpdateProduct)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
