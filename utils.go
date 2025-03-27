// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
package cast

import (
	"math"
	"reflect"
)

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// typeName returns the name of the type T as a string.
func typeName[T any]() string {
	var sample T
	return reflect.TypeOf(sample).String()
}

// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(value any) any {
	if value == nil {
		return nil
	}

	if t := reflect.TypeOf(value); t.Kind() != reflect.Pointer {
		return value
	}

	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// inRange checks if `i` can be safely converted to type `T`.
func inRange[T numeric, I numeric](i I) (T, bool) {
	var zero T
	switch any(zero).(type) {
	case int:
		if int64(i) >= math.MinInt && int64(i) <= math.MaxInt {
			return T(i), true
		}
	case int8:
		if int64(i) >= math.MinInt8 && int64(i) <= math.MaxInt8 {
			return T(i), true
		}
	case int16:
		if int64(i) >= math.MinInt16 && int64(i) <= math.MaxInt16 {
			return T(i), true
		}
	case int32:
		if int64(i) >= math.MinInt32 && int64(i) <= math.MaxInt32 {
			return T(i), true
		}
	case int64:
		return T(i), true
	case uint:
		if i >= 0 && uint64(i) <= math.MaxUint {
			return T(i), true
		}
	case uint8:
		if i >= 0 && uint64(i) <= math.MaxUint8 {
			return T(i), true
		}
	case uint16:
		if i >= 0 && uint64(i) <= math.MaxUint16 {
			return T(i), true
		}
	case uint32:
		if i >= 0 && uint64(i) <= math.MaxUint32 {
			return T(i), true
		}
	case uint64:
		if i >= 0 {
			return T(i), true
		}
	case float32:
		if float64(i) >= -math.MaxFloat32 && float64(i) <= math.MaxFloat32 {
			return T(i), true
		}
	case float64:
		return T(i), true
	}
	return zero, false
}
