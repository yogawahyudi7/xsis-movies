package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func TemplateMessage(field reflect.StructField, err validator.FieldError) (message string) {
	fieldName := field.Tag.Get("json")

	message = fmt.Sprintf("Sorry, Validation failed on field %v with tag '%v'", fieldName, err.Tag())

	return message

}
