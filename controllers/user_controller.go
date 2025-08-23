package controllers

import (
	"golang-simple-rest-api/database"
	"golang-simple-rest-api/helpers"
	"golang-simple-rest-api/payloads"
	"golang-simple-rest-api/services"
	"strconv"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService services.UserService
}

func (controller UserController) CreateUser(c *fiber.Ctx) error {
	payload := payloads.CreateUserPayload{}
	if err := helpers.ValidateBody(&payload, c); err != nil {
		return err
	}

	// if err := c.BodyParser(payload); err != nil {
	// 	log.Println("Error parsing request body: ", err, payload)
	// 	return helpers.Response(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	// }

	tx := database.ClientPostgres.Begin()
	if err := controller.UserService.CreateUser(payload, c, tx); err != nil {
		log.Error("Error creating user: ", err)
		tx.Rollback()
		if helpers.IsDuplicateKeyError(err) {
			return helpers.Response(c, fiber.StatusBadRequest, "Email already exists", err.Error())
		}
		return helpers.Response(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return helpers.Response(c, fiber.StatusInternalServerError, "Failed to create user when save transcation", err.Error())
	}

	return nil
}

func (controller UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Info("get user by id: ", id)
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error("Error parsing ID: ", err)
		return helpers.Response(c, fiber.StatusBadRequest, "Invalid ID", nil)
	}

	if err := controller.UserService.GetUserByID(uint(idUint), database.ClientPostgres, c); err != nil {
		return err
	}

	return nil
}
