package dto

import "github.com/go-playground/validator/v10"

type CreateTaskRequest struct {
	Title string `json:"title" validate:"required,min=2,max=100"`
	Done  bool   `json:"done"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateCreateTask(req CreateTaskRequest) map[string]string {
	err := validate.Struct(req)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		if tag == "required" {
			errors[field] = field + " is required"
		} else if tag == "min" {
			errors[field] = field + " must be at least 2 characters"
		} else if tag == "max" {
			errors[field] = field + " must be less than 100 characters"
		}
	}
	return errors
}
