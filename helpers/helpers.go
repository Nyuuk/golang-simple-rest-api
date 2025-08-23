package helpers

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Payload interface {
	CustomErrorsMessage(validator.ValidationErrors) []map[string]string
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetDSN() string {
	host := GetEnv("DB_HOST", "localhost")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "postgres")
	port := GetEnv("DB_PORT", "5432")
	dbname := GetEnv("DB_NAME", "simple_rest_api")
	timezone := GetEnv("TIMEZONE", "Asia/Jakarta")
	applicationName := GetEnv("APPLICATION_NAME", "simple-rest-api")

	const dsnPattern = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s application_name=%s"
	return fmt.Sprintf(dsnPattern,
		host, user, password, dbname, port, timezone, applicationName)
}

func ValidateBody(payload Payload, c *fiber.Ctx) error {
	log.Debug("Validating request body ", payload)
	validate := validator.New()

	// Parse the request body
	if err := c.BodyParser(&payload); err != nil {
		return Response(c, fiber.StatusBadRequest, "Invalid payload", nil)
	}
	log.Debug("Validating after body parser ", payload)

	// Validate the user struct
	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		data := payload.CustomErrorsMessage(validationErrors)
		log.Debug("Validation errors ", data)
		// log.Debug("Validation errors ", validationErrors)
		return ErrorClient("Invalid payload ", fiber.StatusBadRequest, data)
		// return ErrorClient("Invalid payload ", fiber.StatusBadRequest, validationErrors)
	}

	return nil
}
