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
		log.Error("Error validating request body: ", err)
		if customErr, ok := err.(helpers.Error); ok {
			return helpers.ResponseErrorBadRequest(c, customErr.Message, customErr.Data)
		}
		return err
	}

	// if err := c.BodyParser(payload); err != nil {
	// 	log.Println("Error parsing request body: ", err, payload)
	// 	return helpers.Response(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	// }

	// using transaction
	tx := database.ClientPostgres.Begin()
	if err := controller.UserService.CreateUser(payload, c, tx); err != nil {
		log.Error("Error creating user: ", err)
		tx.Rollback()
		if helpers.IsDuplicateKeyError(err) {
			return helpers.ResponseErrorBadRequest(c, "Email already exists", err)
		}
		return helpers.ResponseErrorInternal(c, err)
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
		return helpers.ResponseErrorBadRequest(c, "Invalid ID", err)
	}

	if err := controller.UserService.GetUserByID(uint(idUint), database.ClientPostgres, c); err != nil {
		log.Error("Error getting user by id: ", err)
		return helpers.ResponseErrorInternal(c, err)
	}

	return nil
}

func (controller UserController) GetAllUsers(c *fiber.Ctx) error {
	log.Info("get all users")
	if err := controller.UserService.GetAllUsers(database.ClientPostgres, c); err != nil {
		log.Error("Error getting all users: ", err)
		return helpers.ResponseErrorInternal(c, err)
	}

	return nil
}

func (controller UserController) DeleteUser(c *fiber.Ctx) error {
	log.Info("delete user")
	id := c.Params("id")
	log.Info("delete user by id: ", id)
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error("Error parsing ID: ", err)
		return helpers.ResponseErrorBadRequest(c, "Invalid ID", err)
	}

	// doen't using transaction
	if err := controller.UserService.DeleteUser(uint(idUint), database.ClientPostgres, c); err != nil {
		log.Error("Error deleting user: ", err)
		return helpers.ResponseErrorInternal(c, err)
	}
	return nil
}

func (controller UserController) UpdateUser(c *fiber.Ctx) error {
	log.Info("update user")

	// // validate manual
	// payload := payloads.UpdateUserPayload{}
	// if err := c.BodyParser(&payload); err != nil {
	// 	log.Error("Error parsing request body: ", err)
	// 	return helpers.ResponseErrorBadRequest(c, "Invalid request body", err)
	// }
	// // validate by payload validate.struct
	// validate := validator.New()
	// if err := validate.Struct(payload); err != nil {
	// 	validationErrors := err.(validator.ValidationErrors)
	// 	log.Debug("Validation errors ", validationErrors)
	// 	return helpers.ResponseErrorBadRequest(c, "Invalid request body", validationErrors)
	// }
	payload := payloads.UpdateUserPayload{}
	if err := helpers.ValidateBody(&payload, c); err != nil {
		log.Error("Error validating request body: ", err)
		if customErr, ok := err.(helpers.Error); ok {
			return helpers.ResponseErrorBadRequest(c, customErr.Message, customErr.Data)
		}
		return err
	}

	// doen't using transaction
	if err := controller.UserService.UpdateUser(payload, database.ClientPostgres, c); err != nil {
		log.Error("Error updating user: ", err)
		return helpers.ResponseErrorInternal(c, err)
	}
	return nil
}
