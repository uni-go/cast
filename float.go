package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

// Float32Provider defines an interface for providing a float32 value with an error.
type Float32Provider interface {
	Float32() (float32, error)
}

// Float32SliceProvider defines an interface for providing a slice of float32 slice with an error.
type Float32SliceProvider interface {
	Float32Slice() ([]float32, error)
}

// Float64Provider defines an interface for providing a float64 value with an error.
type Float64Provider interface {
	Float64() (float64, error)
}

// Float64SliceProvider defines an interface for providing a slice of float64 slice with an error.
type Float64SliceProvider interface {
	Float64Slice() ([]float64, error)
}

// ToFloat converts an interface to a float type (float32 or float64).
func ToFloat[T float32 | float64](value interface{}) (T, error) {
	value = indirect(value)
	title := typeName[T]()
	tError := newTypeError(title)
	oError := newOverflowError(title)

	switch value.(type) {
	case nil:
		return 0, newNilError(title)
	}

	// Handle provider interfaces
	var zero T
	switch any(zero).(type) {
	case float32:
		if val, ok := value.(Float32Provider); ok {
			if v, e := val.Float32(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case float64:
		if val, ok := value.(Float64Provider); ok {
			if v, e := val.Float64(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	}

	// Handle basic types and conversions
	switch val := value.(type) {
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case int, int8, int16, int32, int64:
		if v, ok := inRange[T](reflect.ValueOf(val).Int()); ok {
			return v, nil
		}
		return 0, oError
	case uint, uint8, uint16, uint32, uint64:
		if v, ok := inRange[T](reflect.ValueOf(val).Uint()); ok {
			return v, nil
		}
		return 0, oError
	case float32, float64:
		if v, ok := inRange[T](reflect.ValueOf(val).Float()); ok {
			return v, nil
		}
		return 0, oError
	case string:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			if v, ok := inRange[T](f); ok {
				return v, nil
			}
			return 0, oError
		}

		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			if v, ok := inRange[T](i); ok {
				return v, nil
			}
			return 0, oError
		}

		return 0, tError
	default:
		if f, err := strconv.ParseFloat(fmt.Sprint(val), 64); err == nil {
			if v, ok := inRange[T](f); ok {
				return v, nil
			}
			return 0, oError
		}

		if i, err := strconv.ParseInt(fmt.Sprint(val), 10, 64); err == nil {
			if v, ok := inRange[T](i); ok {
				return v, nil
			}
			return 0, oError
		}

		return 0, tError
	}
}

// ToFloatSlice converts an interface to a slice of float types (float32 or float64).
func ToFloatSlice[T float32 | float64](value interface{}) ([]T, error) {
	value = indirect(value)
	title := "[]" + typeName[T]()
	tError := newTypeError(title)
	oError := newOverflowError(title)

	switch v := value.(type) {
	case nil:
		return nil, newNilError(title)
	case []T:
		return v, nil
	}

	// Handle provider interfaces
	var zero T
	switch any(zero).(type) {
	case float32:
		if val, ok := value.(Float32SliceProvider); ok {
			if v, e := val.Float32Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case float64:
		if val, ok := value.(Float64SliceProvider); ok {
			if v, e := val.Float64Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	}

	// Handle slices and arrays
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		arr := reflect.ValueOf(value)
		res := make([]T, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			f, err := ToFloat[T](arr.Index(i).Interface())
			if err != nil {
				if IsNilError(err) {
					return nil, newNilError(title)
				} else if IsCastError(err) {
					return nil, tError
				} else if IsOverflowError(err) {
					return nil, oError
				} else {
					return nil, fmt.Errorf("%s: %w", title, err)
				}
			}
			res = append(res, f)
		}
		return res, nil
	}

	return nil, newTypeError(title)
}
