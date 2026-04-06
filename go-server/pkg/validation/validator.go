package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())
	return &Validator{validate: v}
}

func (v *Validator) Validate(i any) error {
	if i == nil {
		return errors.New("validation target is nil")
	}
	return v.validate.Struct(i)
}

