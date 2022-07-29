package controllers

import (
	"dumbmerch-api/database"
	"dumbmerch-api/models/entity"
	"dumbmerch-api/models/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserGetAll(ctx *fiber.Ctx) error {
	// get payload token
	userData := ctx.Locals("userId")
	log.Println(userData)

	// data := map[string]string{
	// 	"nama":   "samsul",
	// 	"status": "admin",
	// }
	// log.Println(data)

	var users []entity.User
	result := database.DB.Preload("Profile").Preload("Products").Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(fiber.Map{
		"users": users,
	})
}

func UserGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	return ctx.Status(404).JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}
