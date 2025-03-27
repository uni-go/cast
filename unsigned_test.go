package cast_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToUnsignedUint(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{-1, 0, true},
		{1.5, 1, false},
		{"1", 1, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"invalid", 0, true},
		{true, 1, false},
		{false, 0, false},
		{[]int{1}, 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsigned[uint](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedUint8(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint8
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{-1, 0, true},
		{math.MaxUint8 + 1, 0, true},
		{"255", 255, false},
		{"-1", 0, true},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsigned[uint8](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedUint16(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint16
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{-1, 0, true},
		{math.MaxUint16 + 1, 0, true},
		{"65535", 65535, false},
		{"-1", 0, true},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsigned[uint16](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedUint32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint32
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{-1, 0, true},
		{math.MaxUint32 + 1, 0, true},
		{"4294967295", 4294967295, false},
		{"-1", 0, true},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsigned[uint32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedUint64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint64
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{-1, 0, true},
		{"18446744073709551615", 18446744073709551615, false},
		{"-1", 0, true},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsigned[uint64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint
		err      bool
	}{
		{nil, nil, true},
		{[]uint{1, 2}, []uint{1, 2}, false},
		{[]string{"1", "2"}, []uint{1, 2}, false},
		{[]string{"1", "-1"}, nil, true},
		{[]string{"1", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsignedSlice[uint](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedSliceUint8(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint8
		err      bool
	}{
		{nil, nil, true},
		{[]uint8{1, 2}, []uint8{1, 2}, false},
		{[]int{1, 2}, []uint8{1, 2}, false},
		{[]string{"255", "0"}, []uint8{255, 0}, false},
		{[]string{"255", "-1"}, nil, true},
		{[]string{"255", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxUint8 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsignedSlice[uint8](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedSliceUint16(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint16
		err      bool
	}{
		{nil, nil, true},
		{[]uint16{1, 2}, []uint16{1, 2}, false},
		{[]int{1, 2}, []uint16{1, 2}, false},
		{[]string{"65535", "0"}, []uint16{65535, 0}, false},
		{[]string{"65535", "-1"}, nil, true},
		{[]string{"65535", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxUint16 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsignedSlice[uint16](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedSliceUint32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint32
		err      bool
	}{
		{nil, nil, true},
		{[]uint32{1, 2}, []uint32{1, 2}, false},
		{[]int{1, 2}, []uint32{1, 2}, false},
		{[]string{"4294967295", "0"}, []uint32{4294967295, 0}, false},
		{[]string{"4294967295", "-1"}, nil, true},
		{[]string{"4294967295", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxUint32 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsignedSlice[uint32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToUnsignedSliceUint64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint64
		err      bool
	}{
		{nil, nil, true},
		{[]uint64{1, 2}, []uint64{1, 2}, false},
		{[]int{1, 2}, []uint64{1, 2}, false},
		{[]string{"18446744073709551615", "0"}, []uint64{18446744073709551615, 0}, false},
		{[]string{"18446744073709551615", "-1"}, nil, true},
		{[]string{"18446744073709551615", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToUnsignedSlice[uint64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
