package data

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation '%s' failed on '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate: validate}
}

// Validate the item
// for more detail the returned error can be cast into a validator.ValidationErrors collection
// if ve, ok := err.(validator.ValidationErrors); ok {
//			fmt.Println(ve.Namespace())
//			fmt.Println(ve.Field())
//			fmt.Println(ve.StructNamespace())
//			fmt.Println(ve.StructField())
//			fmt.Println(ve.Tag())
//			fmt.Println(ve.ActualTag())
//			fmt.Println(ve.Kind())
//			fmt.Println(ve.Type())
//			fmt.Println(ve.Value())
//			fmt.Println(ve.Param())
//			fmt.Println()
//	}
func (v *Validation) Validate(i interface{}) ValidationErrors {
	var returnErrs []ValidationError
	if errs, ok := v.validate.Struct(i).(validator.ValidationErrors); ok {
		if errs != nil {
			for _, err := range errs {
				ve := ValidationError{err.(validator.FieldError)}
				returnErrs = append(returnErrs, ve)
			}
		}
	}
	return returnErrs
}

func validateSKU(fl validator.FieldLevel) bool {
	//sku is of format abc-bcd-def
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}
