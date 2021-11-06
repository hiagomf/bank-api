package validations

import (
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// Document é um validador personalizado para CPF e CNPJ
func Document(fl validator.FieldLevel) bool {
	if val, ok := fl.Field().Interface().(string); ok {
		re := regexp.MustCompile(`\d`)
		val = strings.Join(re.FindAllString(val, -1), "")

		switch len(val) {
		case 14:
			val = val[:2] + "." + val[2:5] + "." + val[5:8] + "/" + val[8:12] + "-" + val[12:]
		case 11:
			val = val[:3] + "." + val[3:6] + "." + val[6:9] + "-" + val[9:]
		default:
			return false
		}

		fl.Field().Set(reflect.ValueOf(val))
		return true
	}
	return false
}
