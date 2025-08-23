package services

import (
	"golang-simple-rest-api/entities"
	"golang-simple-rest-api/helpers"
	"golang-simple-rest-api/payloads"
	"golang-simple-rest-api/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func (service UserService) CreateUser(payload payloads.CreateUserPayload, c *fiber.Ctx, tx *gorm.DB) error {
	// user := new(entities.User)
	// if err := c.BodyParser(user); err != nil {
	// 	return err
	// }
	user := entities.User{
		Email: payload.Email,
		Name:  payload.Name,
	}

	log.Info("Creating user: ", user)
	log.Info("Payload: ", payload)

	err := service.userRepo.CreateUser(&user, tx)
	if err != nil {
		return err
	}
	return helpers.Response(c, fiber.StatusCreated, "User created successfully", user)
}

func (service UserService) GetUserByID(id uint, tx *gorm.DB, c *fiber.Ctx) error {
	var user entities.User
	if err := service.userRepo.GetUserByID(id, &user, tx, c); err != nil {
		return err
	}

	return helpers.Response(c, fiber.StatusOK, "User retrieved successfully", user)
}

func (service UserService) GetUserByEmail(email string, user *entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	return service.userRepo.GetUserByEmail(email, user, tx, c)
}

func (service UserService) GetAllUsers(tx *gorm.DB, c *fiber.Ctx) error {
	var users []entities.User
	err := service.userRepo.GetAllUsers(&users, tx, c)
	if err != nil {
		return err
	}
	return helpers.Response(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func (service UserService) UpdateUser(payload payloads.UpdateUserPayload, tx *gorm.DB, c *fiber.Ctx) error {
	user := entities.User{
		ID:    payload.ID,
		Email: payload.Email,
		Name:  payload.Name,
	}
	err := service.userRepo.UpdateUser(&user, tx, c)
	if err != nil {
		return err
	}
	return helpers.Response(c, fiber.StatusOK, "User updated successfully", user)
}

func (service UserService) DeleteUser(id uint, tx *gorm.DB, c *fiber.Ctx) error {
	err := service.userRepo.DeleteUser(id, tx, c)
	if err != nil {
		return err
	}
	return helpers.Response(c, fiber.StatusOK, "User deleted successfully", nil)
}
