package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/sword/api-backend-challenge/model"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err != nil {
		return model.ErrorDiscover(model.BadRequest{DeveloperMessage: err.Error()})
	}
	return nil
}
