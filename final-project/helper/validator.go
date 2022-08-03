package helper

import "github.com/go-playground/validator/v10"

func Validate(p interface{}) error {
	v := validator.New()
	return v.Struct(p)
}
