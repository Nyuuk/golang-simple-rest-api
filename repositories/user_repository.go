package repositories

import (
	"golang-simple-rest-api/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepo struct{}

func (repo UserRepo) CreateUser(user *entities.User, tx *gorm.DB) error {
	return tx.Create(&user).Error
}

func (repo UserRepo) GetUserByID(id uint, user *entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	err := tx.WithContext(c.Context()).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) GetUserByEmail(email string, user *entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	err := tx.WithContext(c.Context()).Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) GetAllUsers(users *[]entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	err := tx.WithContext(c.Context()).Find(&users).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) UpdateUser(user *entities.User, tx *gorm.DB, c *fiber.Ctx) error {
	err := tx.WithContext(c.Context()).Model(&entities.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepo) DeleteUser(id uint, tx *gorm.DB, c *fiber.Ctx) error {
	err := tx.WithContext(c.Context()).Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
