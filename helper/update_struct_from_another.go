package helper

import (
	"reflect"
)

// toUpdate & updateFrom must by not pointer
func UpdateStructFromAnother[T any, R any](toUpdate T, updateFrom R) *T {
	vB := reflect.TypeOf(&updateFrom).Elem()

	// Looping setiap field pada struct Person
	for i := 0; i < vB.NumField(); i++ {
		// Get nama field pada struct Person
		fieldName := vB.Field(i).Name

		if _, ok := reflect.TypeOf(toUpdate).FieldByName(fieldName); ok {
			value := reflect.ValueOf(updateFrom).FieldByName(fieldName)

			reflect.ValueOf(&toUpdate).Elem().FieldByName(fieldName).Set(value)
		}
	}

	return &toUpdate
}
