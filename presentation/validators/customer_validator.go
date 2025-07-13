package validators

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	
	// Custom validation for Thai phone number
	validate.RegisterValidation("thai_phone", validateThaiPhone)
	
	// Custom validation for Thai ID Card
	validate.RegisterValidation("thai_id", validateThaiID)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateThaiPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// Thai phone format: 08xxxxxxxx or 09xxxxxxxx
	matched, _ := regexp.MatchString(`^0[8-9][0-9]{8}$`, phone)
	return matched
}

func validateThaiID(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	// Thai ID card format: 13 digits
	matched, _ := regexp.MatchString(`^[0-9]{13}$`, id)
	return matched
}