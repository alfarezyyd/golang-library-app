package utils

import (
	"github.com/go-playground/validator/v10"
	"golang-library-app/helper"
	"strings"
)

func RegisterCustomValidator() {
	validate := validator.New()
	err := validate.RegisterValidation("enum", EnumGender)
	helper.LogFatalIfError(err)
}

func EnumGender(fl validator.FieldLevel) bool {
	enumString := fl.Param()
	value := fl.Field().String()
	enumSlice := strings.Split(enumString, "-")
	for _, v := range enumSlice {
		if value == v {
			return true
		}
	}
	return false
}
