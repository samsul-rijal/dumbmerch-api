package controllers

import (
	"dumbmerch-api/database"
	"dumbmerch-api/models/entity"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func TransactionGetAll(ctx *fiber.Ctx) error {
	var transactions []entity.Transaction
	result := database.DB.Preload("Product").Preload("Buyer").Preload("Seller").Find(&transactions)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(fiber.Map{
		"transactions": transactions,
	})
}

func TransactionCreate(c *fiber.Ctx) error {
	transaction := new(entity.Transaction)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(transaction); err != nil {
		fmt.Println(transaction)
		return c.Status(400).JSON(fiber.Map{
			"err": "bad request",
		})
	}

	database.DB.Debug().Create(&transaction)
	return c.JSON(fiber.Map{
		"message": "create data successfully",
	})

}
