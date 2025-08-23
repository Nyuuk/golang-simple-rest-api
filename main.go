package main

import (
	"golang-simple-rest-api/controllers"
	"golang-simple-rest-api/database"
	"golang-simple-rest-api/entities"
	"golang-simple-rest-api/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	log.Println("Starting server...")
	// Initialize database connection
	if err := database.PGOpen(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate database
	if err := database.ClientPostgres.AutoMigrate(&entities.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	app := fiber.New()

	app.Get("/healthy", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	userController := controllers.UserController{}
	app.Get("/user/:id", userController.GetUserByID)
	app.Get("/users", userController.GetAllUsers)
	app.Post("/user", userController.CreateUser)
	app.Delete("/user/:id", userController.DeleteUser)
	app.Put("/user/update", userController.UpdateUser)

	port := helpers.GetEnv("PORT_APPLICATION", "3000")
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	log.Printf("Server starting on port %s", port)

}
