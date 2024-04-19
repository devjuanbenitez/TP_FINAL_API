package handlers

import (
	"strconv"

	"github.com/NikSchaefer/go-fiber/database"
	"github.com/NikSchaefer/go-fiber/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateMarca(c *fiber.Ctx) error {
	db := database.DB
	json := new(model.Marca)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "JSON invalido",
		})
	}
	newMarca := model.Marca{
		Nombre: json.Nombre,
	}
	err := db.Create(&newMarca).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func GetMarcaById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}
	marca := model.Marca{}
	query := model.Marca{ID: id}
	err = db.First(&marca, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Marca no encontrado",
		})
	}
	return c.Status(fiber.StatusOK).JSON(marca)
}

func GetAllMarcas(c *fiber.Ctx) error {
	db := database.DB
	var marcas []model.Marca
	err := db.Find(&marcas).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error al recuperar las marcas",
		})
	}
	return c.Status(fiber.StatusOK).JSON(marcas)
}

func UpdateMarca(c *fiber.Ctx) error {
	type UpdateMarcaRequest struct {
		Nombre string `json:"nombre"`
	}
	db := database.DB

	// Parsear el cuerpo de la solicitud JSON
	json := new(UpdateMarcaRequest)
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

	// Buscar la marca por ID
	var found model.Marca
	err = db.First(&found, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Marca no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Actualizar la marca si se ha enviado un nombre
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

func DeleteMarca(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}
	found := model.Marca{}
	query := model.Marca{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Marca no encontrado",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
