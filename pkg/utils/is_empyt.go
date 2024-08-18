package utils

import (
	"errors"
	"reflect"
)

func IsEmpty(s any) error {
	
	val := reflect.ValueOf(s)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.IsZero() && i > 1 {
			return errors.New("todos os campos precisam ser preenchidos")
		}
	}

	return nil
}
