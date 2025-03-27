package cast

// Caster provides methods for type casting and conversion.
type Caster interface {
	// IsNil checks if the value is nil.
	IsNil() bool

	// Interface returns the value as an interface{}.
	Interface() interface{}

	// Slice converts the value to a slice of interface{}.
	Slice() ([]interface{}, error)

	// SliceSafe converts the value to a slice of interface{}, with a fallback on error.
	SliceSafe(fallback []interface{}) []interface{}

	// Unmarshal decodes the value into the provided output using JSON.
	Unmarshal(out interface{}) error

	// Bool converts the value to a bool.
	Bool() (bool, error)

	// BoolSafe converts the value to a bool, with a fallback on error.
	BoolSafe(fallback bool) bool

	// BoolSlice converts the value to a slice of bool.
	BoolSlice() ([]bool, error)

	// BoolSliceSafe converts the value to a slice of bool, with a fallback on error.
	BoolSliceSafe(fallback []bool) []bool

	// Int converts the value to an int.
	Int() (int, error)

	// IntSafe converts the value to an int, with a fallback on error.
	IntSafe(fallback int) int

	// IntSlice converts the value to a slice of int.
	IntSlice() ([]int, error)

	// IntSliceSafe converts the value to a slice of int, with a fallback on error.
	IntSliceSafe(fallback []int) []int

	// Int8 converts the value to an int8.
	Int8() (int8, error)

	// Int8Safe converts the value to an int8, with a fallback on error.
	Int8Safe(fallback int8) int8

	// Int8Slice converts the value to a slice of int8.
	Int8Slice() ([]int8, error)

	// Int8SliceSafe converts the value to a slice of int8, with a fallback on error.
	Int8SliceSafe(fallback []int8) []int8

	// Int16 converts the value to an int16.
	Int16() (int16, error)

	// Int16Safe converts the value to an int16, with a fallback on error.
	Int16Safe(fallback int16) int16

	// Int16Slice converts the value to a slice of int16.
	Int16Slice() ([]int16, error)

	// Int16SliceSafe converts the value to a slice of int16, with a fallback on error.
	Int16SliceSafe(fallback []int16) []int16

	// Int32 converts the value to an int32.
	Int32() (int32, error)

	// Int32Safe converts the value to an int32, with a fallback on error.
	Int32Safe(fallback int32) int32

	// Int32Slice converts the value to a slice of int32.
	Int32Slice() ([]int32, error)

	// Int32SliceSafe converts the value to a slice of int32, with a fallback on error.
	Int32SliceSafe(fallback []int32) []int32

	// Int64 converts the value to an int64.
	Int64() (int64, error)

	// Int64Safe converts the value to an int64, with a fallback on error.
	Int64Safe(fallback int64) int64

	// Int64Slice converts the value to a slice of int64.
	Int64Slice() ([]int64, error)

	// Int64SliceSafe converts the value to a slice of int64, with a fallback on error.
	Int64SliceSafe(fallback []int64) []int64

	// Uint converts the value to a uint.
	Uint() (uint, error)

	// UintSafe converts the value to a uint, with a fallback on error.
	UintSafe(fallback uint) uint

	// UintSlice converts the value to a slice of uint.
	UintSlice() ([]uint, error)

	// UintSliceSafe converts the value to a slice of uint, with a fallback on error.
	UintSliceSafe(fallback []uint) []uint

	// Uint8 converts the value to a uint8.
	Uint8() (uint8, error)

	// Uint8Safe converts the value to a uint8, with a fallback on error.
	Uint8Safe(fallback uint8) uint8

	// Uint8Slice converts the value to a slice of uint8.
	Uint8Slice() ([]uint8, error)

	// Uint8SliceSafe converts the value to a slice of uint8, with a fallback on error.
	Uint8SliceSafe(fallback []uint8) []uint8

	// Uint16 converts the value to a uint16.
	Uint16() (uint16, error)

	// Uint16Safe converts the value to a uint16, with a fallback on error.
	Uint16Safe(fallback uint16) uint16

	// Uint16Slice converts the value to a slice of uint16.
	Uint16Slice() ([]uint16, error)

	// Uint16SliceSafe converts the value to a slice of uint16, with a fallback on error.
	Uint16SliceSafe(fallback []uint16) []uint16

	// Uint32 converts the value to a uint32.
	Uint32() (uint32, error)

	// Uint32Safe converts the value to a uint32, with a fallback on error.
	Uint32Safe(fallback uint32) uint32

	// Uint32Slice converts the value to a slice of uint32.
	Uint32Slice() ([]uint32, error)

	// Uint32SliceSafe converts the value to a slice of uint32, with a fallback on error.
	Uint32SliceSafe(fallback []uint32) []uint32

	// Uint64 converts the value to a uint64.
	Uint64() (uint64, error)

	// Uint64Safe converts the value to a uint64, with a fallback on error.
	Uint64Safe(fallback uint64) uint64

	// Uint64Slice converts the value to a slice of uint64.
	Uint64Slice() ([]uint64, error)

	// Uint64SliceSafe converts the value to a slice of uint64, with a fallback on error.
	Uint64SliceSafe(fallback []uint64) []uint64

	// Float32 converts the value to a float32.
	Float32() (float32, error)

	// Float32Safe converts the value to a float32, with a fallback on error.
	Float32Safe(fallback float32) float32

	// Float32Slice converts the value to a slice of float32.
	Float32Slice() ([]float32, error)

	// Float32SliceSafe converts the value to a slice of float32, with a fallback on error.
	Float32SliceSafe(fallback []float32) []float32

	// Float64 converts the value to a float64.
	Float64() (float64, error)

	// Float64Safe converts the value to a float64, with a fallback on error.
	Float64Safe(fallback float64) float64

	// Float64Slice converts the value to a slice of float64.
	Float64Slice() ([]float64, error)

	// Float64SliceSafe converts the value to a slice of float64, with a fallback on error.
	Float64SliceSafe(fallback []float64) []float64

	// String converts the value to a string.
	String() (string, error)

	// StringSafe converts the value to a string, with a fallback on error.
	StringSafe(fallback string) string

	// StringSlice converts the value to a slice of string.
	StringSlice() ([]string, error)

	// StringSliceSafe converts the value to a slice of string, with a fallback on error.
	StringSliceSafe(fallback []string) []string
}

// NewCaster creates a new Caster instance.
func NewCaster(v interface{}) Caster {
	return &caster{v: indirect(v)}
}
