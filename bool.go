package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

// BoolProvider defines an interface for providing a boolean value with an error.
type BoolProvider interface {
	Bool() (bool, error)
}

// BoolSliceProvider defines an interface for providing a slice of boolean with an error.
type BoolSliceProvider interface {
	BoolSlice() ([]bool, error)
}

// ToBool converts an interface to a bool. Returns an error if the conversion is not possible.
func ToBool(value interface{}) (bool, error) {
	value = indirect(value)
	switch val := value.(type) {
	case nil:
		return false, newNilError("bool")
	case BoolProvider:
		return val.Bool()
	case bool:
		return val, nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(val).Int() != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(val).Uint() != 0, nil
	case float32, float64:
		return reflect.ValueOf(val).Float() != 0, nil
	case string:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return false, newTypeError("bool")
		}
		return v, nil
	default:
		v, err := strconv.ParseBool(fmt.Sprintf("%v", value))
		if err != nil {
			return false, newTypeError("bool")
		}
		return v, nil
	}
}

// ToBoolSlice converts an interface to a slice of bool. Returns an error if the conversion is not possible.
func ToBoolSlice(value interface{}) ([]bool, error) {
	value = indirect(value)
	switch v := value.(type) {
	case nil:
		return nil, newNilError("[]bool")
	case BoolSliceProvider:
		return v.BoolSlice()
	case []bool:
		return v, nil
	}

	// Handle slices or arrays of values
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		arr := reflect.ValueOf(value)
		res := make([]bool, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			b, err := ToBool(arr.Index(i).Interface())
			if err != nil {
				if IsNilError(err) {
					return nil, newNilError("[]bool")
				} else if IsCastError(err) {
					return nil, newTypeError("[]bool")
				} else {
					return nil, fmt.Errorf("[]bool: %w", err)
				}
			}
			res = append(res, b)
		}
		return res, nil
	}

	return nil, newTypeError("[]bool")
}
