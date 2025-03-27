package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

// UintProvider defines an interface for providing a uint value with an error.
type UintProvider interface {
	Uint() (uint, error)
}

// UintSliceProvider defines an interface for providing a slice of uint with an error.
type UintSliceProvider interface {
	UintSlice() ([]uint, error)
}

// Uint8Provider defines an interface for providing a uint8 value with an error.
type Uint8Provider interface {
	Uint8() (uint8, error)
}

// Uint8SliceProvider defines an interface for providing a slice of uint8 with an error.
type Uint8SliceProvider interface {
	Uint8Slice() ([]uint8, error)
}

// Uint16Provider defines an interface for providing a uint16 value with an error.
type Uint16Provider interface {
	Uint16() (uint16, error)
}

// Uint16SliceProvider defines an interface for providing a slice of uint16 with an error.
type Uint16SliceProvider interface {
	Uint16Slice() ([]uint16, error)
}

// Uint32Provider defines an interface for providing a uint32 value with an error.
type Uint32Provider interface {
	Uint32() (uint32, error)
}

// Uint32SliceProvider defines an interface for providing a slice of uint32 with an error.
type Uint32SliceProvider interface {
	Uint32Slice() ([]uint32, error)
}

// Uint64Provider defines an interface for providing a uint64 value with an error.
type Uint64Provider interface {
	Uint64() (uint64, error)
}

// Uint64SliceProvider defines an interface for providing a slice of uint64 with an error.
type Uint64SliceProvider interface {
	Uint64Slice() ([]uint64, error)
}

// ToUnsigned converts an interface to an unsigned integer type (uint, uint8, uint16, uint32, uint64).
func ToUnsigned[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) (T, error) {
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
	case uint:
		if val, ok := value.(UintProvider); ok {
			if v, e := val.Uint(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case uint8:
		if val, ok := value.(Uint8Provider); ok {
			if v, e := val.Uint8(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case uint16:
		if val, ok := value.(Uint16Provider); ok {
			if v, e := val.Uint16(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case uint32:
		if val, ok := value.(Uint32Provider); ok {
			if v, e := val.Uint32(); e != nil {
				return 0, fmt.Errorf("%s: %w", title, e)
			} else {
				return T(v), nil
			}
		}
	case uint64:
		if val, ok := value.(Uint64Provider); ok {
			if v, e := val.Uint64(); e != nil {
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
		if i, err := strconv.ParseUint(val, 10, 64); err == nil {
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
		if i, err := strconv.ParseUint(fmt.Sprint(val), 10, 64); err == nil {
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

// ToUnsignedSlice converts an interface to a slice of unsigned integers (uint, uint8, uint16, uint32, uint64).
func ToUnsignedSlice[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) ([]T, error) {
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
	case uint:
		if val, ok := value.(UintSliceProvider); ok {
			if v, e := val.UintSlice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case uint8:
		if val, ok := value.(Uint8SliceProvider); ok {
			if v, e := val.Uint8Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case uint16:
		if val, ok := value.(Uint16SliceProvider); ok {
			if v, e := val.Uint16Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case uint32:
		if val, ok := value.(Uint32SliceProvider); ok {
			if v, e := val.Uint32Slice(); e != nil {
				return nil, fmt.Errorf("%s: %w", title, e)
			} else {
				result := make([]T, len(v))
				for i, item := range v {
					result[i] = T(item)
				}
				return result, nil
			}
		}
	case uint64:
		if val, ok := value.(Uint64SliceProvider); ok {
			if v, e := val.Uint64Slice(); e != nil {
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
			f, err := ToUnsigned[T](arr.Index(i).Interface())
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
