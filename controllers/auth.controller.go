package controllers

import (
	"dumbmerch-api/database"
	"dumbmerch-api/models/entity"
	"dumbmerch-api/models/request"
	"dumbmerch-api/models/response"
	"dumbmerch-api/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Register(ctx *fiber.Ctx) error {
	user := new(request.RegisterRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword
	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	newProfile := entity.Profile{
		UserID: int(newUser.ID),
	}
	errCreateProfile := database.DB.Create(&newProfile).Error

	if errCreateProfile != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Check email
	var user entity.User
	err := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong email or password!",
		})
	}

	// Check password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong email or password",
		})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	// claims["email"] = user.Email
	// claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	result := response.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"user":  result,
	})

}
