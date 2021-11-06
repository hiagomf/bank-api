package validations

import (
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// StringField é um validador para string
func StringField(fl validator.FieldLevel) bool {
	if val, ok := fl.Field().Interface().(string); ok {
		re := regexp.MustCompile(`\s{2,}`)
		val = strings.TrimSpace(val)
		val = re.ReplaceAllString(val, " ")

		fl.Field().Set(reflect.ValueOf(val))
		return true
	}
	return false
}
