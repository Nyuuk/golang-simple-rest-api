package payloads

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserPayload struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=8"`
}

type UpdateUserPayload struct {
	ID    uint   `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=8"`
}

func (payload *CreateUserPayload) CustomErrorsMessage(validationErrs validator.ValidationErrors) []map[string]string {
	data := []map[string]string{}

	for _, err := range validationErrs {
		switch err.Field() {
		case "Email":
			data = append(data, map[string]string{
				"email": "Email is required and must be a valid email address",
			})
		case "Name":
			data = append(data, map[string]string{
				"name": "Name is required and must be at least 8 characters long",
			})
		}
	}

	return data
}

func (payload *UpdateUserPayload) CustomErrorsMessage(validationErrs validator.ValidationErrors) []map[string]string {
	data := []map[string]string{}

	for _, err := range validationErrs {
		switch err.Field() {
		case "ID":
			data = append(data, map[string]string{
				"id": "ID is required",
			})
		case "Email":
			data = append(data, map[string]string{
				"email": "Email is required and must be a valid email address",
			})
		case "Name":
			data = append(data, map[string]string{
				"name": "Name is required and must be at least 8 characters long",
			})
		}
	}

	return data
}
