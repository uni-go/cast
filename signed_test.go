package cast_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToSignedInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{1.5, 1, false},
		{"1", 1, false},
		{"0", 0, false},
		{"invalid", 0, true},
		{true, 1, false},
		{false, 0, false},
		{[]int{1}, 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToSigned[int](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedInt8(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int8
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{math.MaxInt8 + 1, 0, true},
		{math.MinInt8 - 1, 0, true},
		{"127", 127, false},
		{"-128", -128, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToSigned[int8](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedInt16(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int16
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{math.MaxInt16 + 1, 0, true},
		{math.MinInt16 - 1, 0, true},
		{"32767", 32767, false},
		{"-32768", -32768, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToSigned[int16](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedInt32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int32
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{math.MaxInt32 + 1, 0, true},
		{math.MinInt32 - 1, 0, true},
		{"2147483647", 2147483647, false},
		{"-2147483648", -2147483648, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToSigned[int32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
		err      bool
	}{
		{nil, 0, true},
		{1, 1, false},
		{0, 0, false},
		{"9223372036854775807", 9223372036854775807, false},
		{"-9223372036854775808", -9223372036854775808, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToSigned[int64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int
		err      bool
	}{
		{nil, nil, true},
		{[]int{1, 2}, []int{1, 2}, false},
		{[]string{"1", "2"}, []int{1, 2}, false},
		{[]string{"1", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSignedSlice[int](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedSliceInt8(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int8
		err      bool
	}{
		{nil, nil, true},
		{[]int8{1, 2}, []int8{1, 2}, false},
		{[]int{1, 2}, []int8{1, 2}, false},
		{[]string{"127", "-128"}, []int8{127, -128}, false},
		{[]string{"127", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxInt8 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSignedSlice[int8](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedSliceInt16(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int16
		err      bool
	}{
		{nil, nil, true},
		{[]int16{1, 2}, []int16{1, 2}, false},
		{[]int{1, 2}, []int16{1, 2}, false},
		{[]string{"32767", "-32768"}, []int16{32767, -32768}, false},
		{[]string{"32767", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxInt16 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSignedSlice[int16](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedSliceInt32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int32
		err      bool
	}{
		{nil, nil, true},
		{[]int32{1, 2}, []int32{1, 2}, false},
		{[]int{1, 2}, []int32{1, 2}, false},
		{[]string{"2147483647", "-2147483648"}, []int32{2147483647, -2147483648}, false},
		{[]string{"2147483647", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxInt32 + 1}, nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSignedSlice[int32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToSignedSliceInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int64
		err      bool
	}{
		{nil, nil, true},
		{[]int64{1, 2}, []int64{1, 2}, false},
		{[]int{1, 2}, []int64{1, 2}, false},
		{[]string{"9223372036854775807", "-9223372036854775808"}, []int64{9223372036854775807, -9223372036854775808}, false},
		{[]string{"9223372036854775807", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSignedSlice[int64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
