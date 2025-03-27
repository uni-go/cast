package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

// StringProvider defines an interface for providing a string value with an error.
type StringProvider interface {
	String() (string, error)
}

// StringSliceProvider defines an interface for providing a slice of string with an error.
type StringSliceProvider interface {
	StringSlice() ([]string, error)
}

// ToString converts an interface to a string. Returns an error if the conversion is not possible.
func ToString(value interface{}) (string, error) {
	value = indirect(value)
	switch val := value.(type) {
	case nil:
		return "", newNilError("string")
	case StringProvider:
		return val.String()
	case fmt.Stringer:
		return val.String(), nil
	case bool:
		return strconv.FormatBool(val), nil
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(val).Int(), 10), nil
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(val).Uint(), 10), nil
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(val).Float(), 'f', -1, 64), nil
	case string:
		return val, nil
	default:
		return "", newTypeError("string")
	}
}

// ToStringSlice converts an interface to a slice of string. Returns an error if the conversion is not possible.
func ToStringSlice(value interface{}) ([]string, error) {
	value = indirect(value)
	switch v := value.(type) {
	case nil:
		return nil, newNilError("[]string")
	case StringSliceProvider:
		return v.StringSlice()
	case []string:
		return v, nil
	}

	// Handle slices or arrays of values
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		arr := reflect.ValueOf(value)
		res := make([]string, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			s, err := ToString(arr.Index(i).Interface())
			if err != nil {
				if IsNilError(err) {
					return nil, newNilError("[]string")
				} else if IsCastError(err) {
					return nil, newTypeError("[]string")
				} else {
					return nil, fmt.Errorf("[]string: %w", err)
				}
			}
			res = append(res, s)
		}
		return res, nil
	}

	return nil, newTypeError("[]string")
}
