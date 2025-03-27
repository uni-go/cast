package cast

import (
	"encoding/json"
	"fmt"
)

type caster struct {
	v interface{}
}

func (c caster) IsNil() bool {
	return c.v == nil
}

func (c caster) Interface() interface{} {
	return c.v
}

func (c caster) Slice() ([]interface{}, error) {
	return ToSlice(c.v)
}

func (c caster) SliceSafe(f []interface{}) []interface{} {
	if v, err := ToSlice(c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Unmarshal(out interface{}) error {
	// Try direct unmarshal
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", c.v)), out)
	if err == nil {
		return nil
	}

	// Try marshal and unmarshal
	bytes, err := json.Marshal(c.v)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}

func (c caster) Bool() (bool, error) {
	return ToBool(c.v)
}

func (c caster) BoolSafe(f bool) bool {
	if v, err := ToBool(c.v); err == nil {
		return v
	}

	return f
}

func (c caster) BoolSlice() ([]bool, error) {
	return ToBoolSlice(c.v)
}

func (c caster) BoolSliceSafe(f []bool) []bool {
	if v, err := ToBoolSlice(c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int() (int, error) {
	return ToSigned[int](c.v)
}

func (c caster) IntSafe(f int) int {
	if v, err := ToSigned[int](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) IntSlice() ([]int, error) {
	return ToSignedSlice[int](c.v)
}

func (c caster) IntSliceSafe(f []int) []int {
	if v, err := ToSignedSlice[int](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int8() (int8, error) {
	return ToSigned[int8](c.v)
}

func (c caster) Int8Safe(f int8) int8 {
	if v, err := ToSigned[int8](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int8Slice() ([]int8, error) {
	return ToSignedSlice[int8](c.v)
}

func (c caster) Int8SliceSafe(f []int8) []int8 {
	if v, err := ToSignedSlice[int8](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int16() (int16, error) {
	return ToSigned[int16](c.v)
}

func (c caster) Int16Safe(f int16) int16 {
	if v, err := ToSigned[int16](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int16Slice() ([]int16, error) {
	return ToSignedSlice[int16](c.v)
}

func (c caster) Int16SliceSafe(f []int16) []int16 {
	if v, err := ToSignedSlice[int16](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int32() (int32, error) {
	return ToSigned[int32](c.v)
}

func (c caster) Int32Safe(f int32) int32 {
	if v, err := ToSigned[int32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int32Slice() ([]int32, error) {
	return ToSignedSlice[int32](c.v)
}

func (c caster) Int32SliceSafe(f []int32) []int32 {
	if v, err := ToSignedSlice[int32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int64() (int64, error) {
	return ToSigned[int64](c.v)
}

func (c caster) Int64Safe(f int64) int64 {
	if v, err := ToSigned[int64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Int64Slice() ([]int64, error) {
	return ToSignedSlice[int64](c.v)
}

func (c caster) Int64SliceSafe(f []int64) []int64 {
	if v, err := ToSignedSlice[int64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint() (uint, error) {
	return ToUnsigned[uint](c.v)
}

func (c caster) UintSafe(f uint) uint {
	if v, err := ToUnsigned[uint](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) UintSlice() ([]uint, error) {
	return ToUnsignedSlice[uint](c.v)
}

func (c caster) UintSliceSafe(f []uint) []uint {
	if v, err := ToUnsignedSlice[uint](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint8() (uint8, error) {
	return ToUnsigned[uint8](c.v)
}

func (c caster) Uint8Safe(f uint8) uint8 {
	if v, err := ToUnsigned[uint8](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint8Slice() ([]uint8, error) {
	return ToUnsignedSlice[uint8](c.v)
}

func (c caster) Uint8SliceSafe(f []uint8) []uint8 {
	if v, err := ToUnsignedSlice[uint8](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint16() (uint16, error) {
	return ToUnsigned[uint16](c.v)
}

func (c caster) Uint16Safe(f uint16) uint16 {
	if v, err := ToUnsigned[uint16](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint16Slice() ([]uint16, error) {
	return ToUnsignedSlice[uint16](c.v)
}

func (c caster) Uint16SliceSafe(f []uint16) []uint16 {
	if v, err := ToUnsignedSlice[uint16](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint32() (uint32, error) {
	return ToUnsigned[uint32](c.v)
}

func (c caster) Uint32Safe(f uint32) uint32 {
	if v, err := ToUnsigned[uint32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint32Slice() ([]uint32, error) {
	return ToUnsignedSlice[uint32](c.v)
}

func (c caster) Uint32SliceSafe(f []uint32) []uint32 {
	if v, err := ToUnsignedSlice[uint32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint64() (uint64, error) {
	return ToUnsigned[uint64](c.v)
}

func (c caster) Uint64Safe(f uint64) uint64 {
	if v, err := ToUnsigned[uint64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Uint64Slice() ([]uint64, error) {
	return ToUnsignedSlice[uint64](c.v)
}

func (c caster) Uint64SliceSafe(f []uint64) []uint64 {
	if v, err := ToUnsignedSlice[uint64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Float32() (float32, error) {
	return ToFloat[float32](c.v)
}

func (c caster) Float32Safe(f float32) float32 {
	if v, err := ToFloat[float32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Float32Slice() ([]float32, error) {
	return ToFloatSlice[float32](c.v)
}

func (c caster) Float32SliceSafe(f []float32) []float32 {
	if v, err := ToFloatSlice[float32](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Float64() (float64, error) {
	return ToFloat[float64](c.v)
}

func (c caster) Float64Safe(f float64) float64 {
	if v, err := ToFloat[float64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) Float64Slice() ([]float64, error) {
	return ToFloatSlice[float64](c.v)
}

func (c caster) Float64SliceSafe(f []float64) []float64 {
	if v, err := ToFloatSlice[float64](c.v); err == nil {
		return v
	}

	return f
}

func (c caster) String() (string, error) {
	return ToString(c.v)
}
func (c caster) StringSafe(f string) string {
	if v, err := ToString(c.v); err == nil {
		return v
	}

	return f
}

func (c caster) StringSlice() ([]string, error) {
	return ToStringSlice(c.v)
}

func (c caster) StringSliceSafe(f []string) []string {
	if v, err := ToStringSlice(c.v); err == nil {
		return v
	}

	return f
}
