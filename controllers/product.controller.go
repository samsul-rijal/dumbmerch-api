package controllers

import (
	"dumbmerch-api/database"
	"dumbmerch-api/models/entity"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ProductGetAll(ctx *fiber.Ctx) error {
	var products []entity.Product
	result := database.DB.Preload("User").Find(&products)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(fiber.Map{
		"products": products,
	})
}

func ProductGetById(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")

	var product entity.Product
	err := database.DB.First(&product, "id = ?", productId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	productResponse := entity.Product{
		ID:    product.ID,
		Name:  product.Name,
		Desc:  product.Desc,
		Image: product.Image,
		Qty:   product.Qty,
	}

	return ctx.Status(404).JSON(fiber.Map{
		"message": "success",
		"data":    productResponse,
	})
}

func ProductCreate(c *fiber.Ctx) error {
	product := new(entity.Product)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(product); err != nil {
		fmt.Println(product)
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// Handle file
	file, errFile := c.FormFile("image")
	if errFile != nil {
		log.Println("Error File :", errFile)
	}

	filename := file.Filename

	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))

	if errSaveFile != nil {
		log.Println("Fail to store file")
	}

	newProduct := entity.Product{
		Name:   product.Name,
		Desc:   product.Desc,
		Price:  product.Price,
		Image:  filename,
		Qty:    product.Qty,
		UserID: product.UserID,
	}

	database.DB.Debug().Create(&newProduct)
	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"product": newProduct,
	})

}
