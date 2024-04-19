package handlers

import (
	"strconv"

	"github.com/NikSchaefer/go-fiber/database"
	"github.com/NikSchaefer/go-fiber/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCategory(c *fiber.Ctx) error {
	db := database.DB
	json := new(model.Categoria)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "JSON invalido",
		})
	}
	newCategory := model.Categoria{
		Nombre: json.Nombre,
	}
	err := db.Create(&newCategory).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}
	category := model.Categoria{}
	query := model.Categoria{ID: id}
	err = db.First(&category, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Producto no encontrado",
		})
	}
	return c.Status(fiber.StatusOK).JSON(category)
}

func GetAllCategories(c *fiber.Ctx) error {
	db := database.DB
	var categories []model.Categoria
	err := db.Find(&categories).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error al recuperar las categorías",
		})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

func UpdateCategory(c *fiber.Ctx) error {
	type UpdateCategoryRequest struct {
		Nombre string `json:"nombre"`
	}
	db := database.DB

	// Parsear el cuerpo de la solicitud JSON
	json := new(UpdateCategoryRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "JSON invalido",
		})
	}

	// Obtener el parámetro "id" de la URL
	param := c.Params("id")

	// Convertir el parámetro "id" de string a int
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}

	// Buscar la categoría por ID
	var found model.Categoria
	err = db.First(&found, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Categoria no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Actualizar la categoría si se ha enviado un nombre
	if json.Nombre != "" {
		found.Nombre = json.Nombre
	}

	// Guardar la categoría actualizada
	db.Save(&found)

	// Devolver una respuesta de éxito
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    found,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}
	found := model.Categoria{}
	query := model.Categoria{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Categoria no encontrado",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
