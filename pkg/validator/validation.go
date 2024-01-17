package validator

import (
	validationLib "github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"regexp"
)

func ValidateEntity(entityToValidate interface{}) error {
	validator := validationLib.New(validationLib.WithRequiredStructEnabled())
	err := validator.RegisterValidation("is_phone", PhoneValidation)
	if err != nil {
		logrus.Fatalf("Cannot create custom validation function %s", err)
	}
	err = validator.Struct(entityToValidate)
	if err != nil {
		return err
	}
	return nil
}

func PhoneValidation(field validationLib.FieldLevel) bool {
	value := field.Field().String()
	pattern := "^\\+[0-9]{11}$"
	matched, err := regexp.Match(pattern, []byte(value))
	if err != nil {
		logrus.Errorf("Error during validation has occurred: %s", err)
	}
	return matched
}
