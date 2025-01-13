package controllers

import (
	"rbac-go/config"
	"rbac-go/models"
	"rbac-go/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct{
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	if !utils.ComparePassword(user.Password, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token, err := utils.GenerateToken(user.ID, user.Roles[0].Name)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"token": token,
	})
}

func Register(c*fiber.Ctx) error{
	type RegisterInput struct {
		UserName string `json:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Role     string `json:"role" validate:"required"`
	}
	

	var input RegisterInput
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


	if err := config.DB.Where("email = ?", input.Email).First(&models.User{}).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email already exists",
		})
	}

	if err := config.DB.Where("username = ?", input.UserName).First(&models.User{}).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username already exists",
		})
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	var role models.Role
	if err := config.DB.Where("name = ?", input.Role).First(&role).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "role not found",
		})
	}

	user := models.User{
		Username: input.UserName,
		Email:    input.Email,
		Password: hashedPassword,
		Roles:    []models.Role{role},
	}
	

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "User Registered Successfully",
		"user": user,
	})
}