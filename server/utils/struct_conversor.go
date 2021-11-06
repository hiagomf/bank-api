package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

const (
	tagName = "conversorTag"
)

// ModelStruct é uma estrutura base que é implementada pelas models
// na camada de dados para possibilitar a conversão entre estruturas diferentes
type ModelStruct interface {
	ConvertFromStruct(interface{}) error
	ConvertToStruct(interface{}) error
}

// ConvertStruct converte os dados da estrutura "in" para a estrutura "out"
// desde que eles possuam a mesma tag "codinome"
func ConvertStruct(in interface{}, out interface{}) error {
	elemSrc := reflect.ValueOf(in).Elem()
	elemDst := reflect.ValueOf(out).Elem()

	if elemSrc.Kind() != reflect.Struct || elemDst.Kind() != reflect.Struct {
		return errors.New("não é uma estrutura")
	}

	if elemSrc.NumField() == 0 || elemDst.NumField() == 0 {
		return errors.New("os campos não estão disponíveis")
	}

	for s := 0; s < elemSrc.NumField(); s++ {
		srcField := elemSrc.Type().Field(s)

		srcKey := srcField.Tag.Get(tagName)
		if srcKey == "" {
			continue
		}
		if elemSrc.Field(s).IsNil() {
			continue
		}

		for d := 0; d < elemDst.NumField(); d++ {
			dstField := elemDst.Type().Field(d)

			dstKey := dstField.Tag.Get(tagName)
			if dstKey == "" || dstKey != srcKey {
				continue
			}

			if !elemDst.Field(d).CanSet() {
				return errors.New("não é possível atribuir valor ao destino")
			}

			if srcField.Type != dstField.Type {
				var (
					tSrc, tDst string
				)

				if srcField.Type.String()[0] == '*' {
					tSrc = srcField.Type.Elem().String()
				} else {
					tSrc = srcField.Type.String()
				}

				if dstField.Type.String()[0] == '*' {
					tDst = dstField.Type.Elem().String()
				} else {
					tDst = dstField.Type.String()
				}

				if dstField.Type.String() == "*time.Time" {
					tDst = dstField.Type.String()
				}

				val, err := converta(elemSrc.Field(s).Interface(), tSrc, tDst)
				if err != nil {
					return fmt.Errorf("não foi possível converter campo %s: %s", srcField.Name, err.Error())
				}

				v := reflect.ValueOf(val)
				elemDst.Field(d).Set(v)
			} else {
				elemDst.Field(d).Set(elemSrc.Field(s))
			}
		}
	}

	return nil
}

func converta(value interface{}, from, to string) (output interface{}, err error) {
	if from == "string" {
		if to == "*time.Time" {
			var val string
			if v, ok := value.(*string); ok {
				if v == nil {
					return nil, nil
				}
				val = *v
			} else if v, ok := value.(string); ok {
				val = v
			} else {
				return nil, nil
			}
			return ParseDataTempo(val)
		}
	}
	return nil, errors.New("valor não suportado")
}

// ParseDataTempo tentar dar parse em uma string para um tipo time.Time
func ParseDataTempo(value string) (*time.Time, error) {
	// assumimos que timezone que deve ser utilizada é a America/Fortaleza
	fortaleza, err := time.LoadLocation("America/Fortaleza")
	if err != nil {
		return nil, err
	}

	// https://golang.org/pkg/time/#pkg-constants
	// Mon Jan 2 15:04:05 MST 2006
	layout := `02/01/2006 15:04:05`

	switch {
	case regexp.MustCompile(`^(\d{0,2})\/(\d{0,2})\/(\d{0,4})\s(\d{0,2}):(\d{0,2}):(\d{0,2})\.?\d{0,} (-|\+)(\d{2}):(\d{2})$`).Match([]byte(value)):
		// the string contains a timezone indicator with hour and minutes separator
		layout += " -07:00"
	case regexp.MustCompile(`^(\d{0,2})\/(\d{0,2})\/(\d{0,4})\s(\d{0,2}):(\d{0,2}):(\d{0,2})\.?\d{0,} (-|\+)(\d{2})(\d{2})$`).Match([]byte(value)):
		// the string contains a timezone indicator
		layout += " -0700"
	case regexp.MustCompile(`^(\d{0,2})\/(\d{0,2})\/(\d{0,4})\s(\d{0,2}):(\d{0,2}):(\d{0,2})\.?\d{0,} (-|\+)(\d{2})$`).Match([]byte(value)):
		// the string contains a half-assed timezone indicator
		layout += " -07"
	case regexp.MustCompile(`^(\d{0,2})\/(\d{0,2})\/(\d{0,4})\s(\d{0,2}):(\d{0,2}):(\d{0,2})\.?\d{0,}$`).Match([]byte(value)):
		// the string matches the whole format no changes are necessary
	case regexp.MustCompile(`^(\d{0,2})\/(\d{0,2})\/(\d{0,4})$`).Match([]byte(value)):
		// the string doesn't contain time information so we append to it
		value += " 00:00:00"
	default:
		return nil, fmt.Errorf("o layout data não pode ser identificado `%s`", value)
	}

	t, err := time.ParseInLocation(layout, value, fortaleza)
	return &t, err
}
