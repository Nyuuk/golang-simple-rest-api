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

func (service UserService) GetAllUsers(users *[]entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	return service.userRepo.GetAllUsers(users, tx, c)
}

func (service UserService) UpdateUser(user *entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	return service.userRepo.UpdateUser(user, tx, c)
}

func (service UserService) DeleteUser(id uint, tx *gorm.DB, c *fiber.Ctx) error {
	return service.userRepo.DeleteUser(id, tx, c)
}
