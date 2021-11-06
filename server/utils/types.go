package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

// CriadorID é um tipo usado na documentação
type CriadorID struct {
	ID int64 `json:"id"`
}

// NoContent define o a model para respostas vazias
type NoContent struct{}

// JSONB representa um tipo JSOB do postgres
type JSONB map[string]interface{}

// Scan implementa a interface sql.Scaner
func (j *JSONB) Scan(src interface{}) error {
	val, ok := src.([]byte)

	if !ok {
		return errors.New("pointeiro de tipo JSONB é inválido")
	}

	if err := json.Unmarshal(val, j); err != nil {
		return err
	}

	return nil
}

// Value implementa driver.Value
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	s, err := Stringify(j)
	if err != nil {
		return nil, err
	}

	return *s, nil
}

// MarshalLogObject implementa a interface do zap
func (j JSONB) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if j == nil {
		return
	}

	for k, v := range j {
		switch v := v.(type) {
		case bool:
			enc.AddBool(k, v)
		case int:
			enc.AddInt(k, v)
		case int32:
			enc.AddInt32(k, v)
		case int64:
			enc.AddInt64(k, v)
		case float32:
			enc.AddFloat32(k, v)
		case float64:
			enc.AddFloat64(k, v)
		case string:
			enc.AddString(k, v)
		case time.Time:
			enc.AddTime(k, v)
		case time.Duration:
			enc.AddDuration(k, v)
		default:
			if err = enc.AddReflected(k, v); err != nil {
				return
			}
		}
	}

	return
}

// JSONBA representa um array de JSONB
type JSONBA []JSONB

// MarshalLogArray implementa a interface do zap
func (j JSONBA) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	if j == nil {
		return
	}

	for i := range j {
		if err = enc.AppendObject(j[i]); err != nil {
			return
		}
	}

	return
}

// Scan implementa a interface do sql.Scaner
func (j *JSONBA) Scan(src interface{}) error {
	val, ok := src.([]byte)

	if !ok {
		return errors.New("ponteiro para tipo JSONBA é inválido")
	}

	if err := json.Unmarshal(val, j); err != nil {
		return err
	}

	return nil
}

// Value implementa a interface de driver.Value
func (j JSONBA) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	var s []string = make([]string, len(j))

	for i := range j {
		o, err := Stringify(j[i])
		if err != nil {
			return nil, err
		}
		s[i] = *o
	}

	return "[" + strings.Join(s, ",") + "]", nil
}
