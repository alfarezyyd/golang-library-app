package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
	"time"
)

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", enumGender)
		if err != nil {
			log.Fatal(err)
		}
		err = v.RegisterValidation("birth_date", birthDate)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var enumGender validator.Func = func(fl validator.FieldLevel) bool {
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

var birthDate validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, err := time.Parse("02-01-2006", value)
	if err != nil {
		return false
	}
	return true
}
