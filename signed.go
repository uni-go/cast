package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

// IntProvider defines an interface for providing a int value with an error.
type IntProvider interface {
	Int() (int, error)
}

// IntSliceProvider defines an interface for providing a slice of int with an error.
type IntSliceProvider interface {
	IntSlice() ([]int, error)
}

// Int8Provider defines an interface for providing a int8 value with an error.
type Int8Provider interface {
	Int8() (int8, error)
}

// Int8SliceProvider defines an interface for providing a slice of int8 with an error.
type Int8SliceProvider interface {
	Int8Slice() ([]int8, error)
}

// Int16Provider defines an interface for providing a int16 value with an error.
type Int16Provider interface {
	Int16() (int16, error)
}

// Int16SliceProvider defines an interface for providing a slice of int16 with an error.
type Int16SliceProvider interface {
	Int16Slice() ([]int16, error)
}

// Int32Provider defines an interface for providing a int32 value with an error.
type Int32Provider interface {
	Int32() (int32, error)
}

// Int32SliceProvider defines an interface for providing a slice of int32 with an error.
type Int32SliceProvider interface {
	Int32Slice() ([]int32, error)
}

// Int64Provider defines an interface for providing a int64 value with an error..
type Int64Provider interface {
	Int64() (int64, error)
}

// Int64SliceProvider defines an interface for providing a slice of int64 with an error.
type Int64SliceProvider interface {
	Int64Slice() ([]int64, error)
}

// ToSigned converts an interface to a signed integer type (int, int8, int16, int32, int64).
func ToSigned[T int | int8 | int16 | int32 | int64](value interface{}) (T, error) {
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
	case int:
		if val, ok := value.(IntProvider); ok {
			if v, e := val.Int(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case int8:
		if val, ok := value.(Int8Provider); ok {
			if v, e := val.Int8(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case int16:
		if val, ok := value.(Int16Provider); ok {
			if v, e := val.Int16(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case int32:
		if val, ok := value.(Int32Provider); ok {
			if v, e := val.Int32(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case int64:
		if val, ok := value.(Int64Provider); ok {
			if v, e := val.Int64(); e != nil {
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
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			if v, ok := inRange[T](i); ok {
				return v, nil
			}
			return 0, oError
		}

		if f, err := strconv.ParseFloat(val, 64); err == nil {
			if v, ok := inRange[T](f); ok {
				return v, nil
			}
			return 0, oError
		}

		return 0, tError
	default:
		if i, err := strconv.ParseInt(fmt.Sprint(val), 10, 64); err == nil {
			if v, ok := inRange[T](i); ok {
				return v, nil
			}
			return 0, oError
		}

		if f, err := strconv.ParseFloat(fmt.Sprint(val), 64); err == nil {
			if v, ok := inRange[T](f); ok {
				return v, nil
			}
			return 0, oError
		}

		return 0, tError
	}
}

// ToSignedSlice converts an interface to a slice of signed integers (int, int8, int16, int32, int64).
func ToSignedSlice[T int | int8 | int16 | int32 | int64](value interface{}) ([]T, error) {
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
	case int:
		if val, ok := value.(IntSliceProvider); ok {
			if v, e := val.IntSlice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case int8:
		if val, ok := value.(Int8SliceProvider); ok {
			if v, e := val.Int8Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case int16:
		if val, ok := value.(Int16SliceProvider); ok {
			if v, e := val.Int16Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case int32:
		if val, ok := value.(Int32SliceProvider); ok {
			if v, e := val.Int32Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case int64:
		if val, ok := value.(Int64SliceProvider); ok {
			if v, e := val.Int64Slice(); e != nil {
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
			f, err := ToSigned[T](arr.Index(i).Interface())
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
