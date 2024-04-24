package util

import (
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate implements echo's validator interface using the go-playground validator.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
