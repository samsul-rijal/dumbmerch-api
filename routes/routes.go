package routes

import (
	"dumbmerch-api/controllers"
	"dumbmerch-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router) {
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)

	r.Get("/users", middleware.Auth, controllers.UserGetAll)
	r.Get("/user/:id", middleware.Auth, controllers.UserGetById)

	r.Get("/products", controllers.ProductGetAll)
	r.Get("/product/:id", middleware.Auth, controllers.ProductGetById)
	r.Post("/product", middleware.Auth, controllers.ProductCreate)

	r.Get("/transactions", controllers.TransactionGetAll)
	r.Post("/transaction", controllers.TransactionCreate)
}
