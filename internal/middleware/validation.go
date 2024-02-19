package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/rulanugrh/lysithea/internal/entity/web"
)

type ValidationInterface interface {
	Validate(data interface{}) error
	ValidationMessage(err error) error
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() ValidationInterface {
	return &Validation{
		validate: validator.New(),
	}
}

func (v *Validation) Validate(data interface{}) error {
	err := v.validate.Struct(data)
	if err != nil {
		return err
	}

	return nil
}
func (v *Validation) ValidationMessage(err error) error {
	var msg string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", e.Field())
		case "email":
			msg = fmt.Sprintf("%s data must be email format", e.Field())
		case "min":
			msg = fmt.Sprintf("%s is to short", e.Field())
		}
	}

	return web.StatusBadRequest(msg)
}
