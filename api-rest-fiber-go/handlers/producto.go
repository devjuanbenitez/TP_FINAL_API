package handlers

import (
	"strconv"

	"github.com/NikSchaefer/go-fiber/database"
	"github.com/NikSchaefer/go-fiber/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB

	// Parsear el cuerpo de la solicitud JSON
	productRequest := struct {
		Nombre      string  `json:"nombre"`
		Precio      float64 `json:"precio"`
		Stock       int     `json:"stock"`
		MarcaID     int     `json:"marca"`
		CategoriaID int     `json:"categoria"`
	}{}
	if err := c.BodyParser(&productRequest); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "JSON invalido",
		})
	}

	// Verificar si existe la marca con el ID proporcionado
	var marca model.Marca
	err := db.First(&marca, productRequest.MarcaID).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Marca no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Verificar si existe la categoría con el ID proporcionado
	var categoria model.Categoria
	err = db.First(&categoria, productRequest.CategoriaID).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Categoria no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Crear una instancia del modelo Producto
	product := model.Producto{
		Nombre:      productRequest.Nombre,
		Precio:      productRequest.Precio,
		Stock:       productRequest.Stock,
		MarcaID:     productRequest.MarcaID,
		CategoriaID: productRequest.CategoriaID,
	}

	// Crear el nuevo producto en la base de datos
	err = db.Create(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Devolver una respuesta de éxito
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    product,
	})
}

func GetProductById(c *fiber.Ctx) error {
	db := database.DB

	// Obtener el parámetro "id" de la URL
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}

	// Crear una instancia de Producto para almacenar el resultado de la consulta
	var product model.Producto

	// Realizar la consulta para encontrar el producto con el ID especificado
	err = db.First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Producto no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Devolver el producto encontrado como JSON
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    product,
	})
}
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Producto
	err := db.Find(&products).Error
	if err != nil {
		// Manejar el error si la consulta falla
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error al recuperar los productos",
		})
	}

	// Devuelve los productos como una respuesta JSON
	return c.Status(fiber.StatusOK).JSON(products)
}

func UpdateProduct(c *fiber.Ctx) error {
	// Define la estructura para los datos de la solicitud
	type UpdateProductRequest struct {
		Nombre      string  `json:"nombre"`
		CategoriaID int     `json:"categoria"`
		MarcaID     int     `json:"marca"`
		Precio      float64 `json:"precio"`
		Stock       int     `json:"stock"`
	}

	db := database.DB

	// Parsear el cuerpo de la solicitud JSON
	json := new(UpdateProductRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "JSON invalido",
		})
	}

	// Obtener el parámetro "id" de la URL
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}

	// Verificar si los IDs de Categoria y Marca existen en la base de datos
	var categoria model.Categoria
	var marca model.Marca

	// Verificar la existencia de la categoría
	if json.CategoriaID != 0 {
		err = db.First(&categoria, json.CategoriaID).Error
		if err == gorm.ErrRecordNotFound {
			return c.JSON(fiber.Map{
				"code":    400,
				"message": "Categoria no encontrado",
			})
		} else if err != nil {
			return c.JSON(fiber.Map{
				"code":    500,
				"message": "Error interno del servidor",
			})
		}
	}
	// Verificar la existencia de la marca
	if json.MarcaID != 0 {
		err = db.First(&marca, json.MarcaID).Error
		if err == gorm.ErrRecordNotFound {
			return c.JSON(fiber.Map{
				"code":    400,
				"message": "Marca no encontrado",
			})
		} else if err != nil {
			return c.JSON(fiber.Map{
				"code":    500,
				"message": "Error interno del servidor",
			})
		}
	}
	// Buscar el producto por ID
	var product model.Producto
	err = db.First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Producto no encontrado",
		})
	} else if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}
	// Actualizar el producto con los datos recibidos
	if json.Nombre != "" {
		product.Nombre = json.Nombre
	}
	if json.CategoriaID != 0 {
		product.CategoriaID = json.CategoriaID
	}
	if json.MarcaID != 0 {
		product.MarcaID = json.MarcaID
	}
	if json.Precio != 0 {
		product.Precio = json.Precio
	}
	if json.Stock != 0 {
		product.Stock = json.Stock
	}
	db.Save(&product)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB

	// Obtener el parámetro "id" de la URL
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Formato de ID invalido",
		})
	}

	// Buscar el producto por ID
	var product model.Producto
	err = db.First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Producto no encontrado",
		})
	}

	// Eliminar el producto
	err = db.Delete(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Error interno del servidor",
		})
	}

	// Devolver una respuesta de éxito
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
