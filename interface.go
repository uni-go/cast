package cast

import (
	"reflect"
)

// SliceProvider defines an interface for providing a slice of interface{} with an error.
type SliceProvider interface {
	Slice() ([]interface{}, error)
}

// ToSlice converts an interface to a slice of interface{}. Returns an error if the conversion is not possible.
func ToSlice(value interface{}) ([]interface{}, error) {
	value = indirect(value)
	switch v := value.(type) {
	case nil:
		return nil, newNilError("[]interface{}")
	case SliceProvider:
		return v.Slice()
	}

	// Handle slices or arrays of values
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		arr := reflect.ValueOf(value)
		res := make([]interface{}, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			res = append(res, arr.Index(i).Interface())
		}
		return res, nil
	}

	return nil, newTypeError("[]interface{}")
}
