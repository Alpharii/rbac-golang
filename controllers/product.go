package controllers

import (
	"rbac-go/config"
	"rbac-go/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreateProduct(c*fiber.Ctx) error{
	type ProductInput struct {
		Name        string  `json:"name" validate:"required,max=100"`
		Description string  `json:"description" validate:"required,max=255"`
		Price       float64 `json:"price" validate:"required,gt=0"`
	}

	var input ProductInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
			"details": err.Error(),
		})
	}

	product := models.Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price

	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "product created",
		"product": product,
	})
}

func GetAllProducts(c*fiber.Ctx) error{
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get products",
		})
	}
	config.DB.Find(&products)
	if(len(products) == 0){
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "no products found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "products found",
		"products": products,
	})
}

func GetSingleProduct(c*fiber.Ctx) error{
	id := c.Params("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get product",
		})
	}
	return c.JSON(fiber.Map{
		"message": "product found",
		"product": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	type ProductInput struct {
		Name        *string  `json:"name" validate:"omitempty,max=100"`
		Description *string  `json:"description" validate:"omitempty,max=255"`
		Price       *float64 `json:"price" validate:"omitempty,gt=0"`
	}

	var input ProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid input",
			"details": err.Error(),
		})
	}

	var product models.Product

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}

	if err := config.DB.Where("id = ?", id).Updates(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product with id " + id + " updated",
		"product": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete product",
		})
	}
	return c.JSON(fiber.Map{
		"message": "product with id " + id + " deleted",
		"product": product,
	})
}