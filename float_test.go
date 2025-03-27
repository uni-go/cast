package cast_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToFloat32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float32
		err      bool
	}{
		{nil, 0, true},
		{1, 1.0, false},
		{0, 0.0, false},
		{1.5, 1.5, false},
		{"1.5", 1.5, false},
		{"0", 0.0, false},
		{"invalid", 0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{[]int{1}, 0, true},
		{math.MaxFloat64, 0, true},     // Overflow case
		{math.MaxFloat32 * 2, 0, true}, // Overflow case
	}

	for _, test := range tests {
		result, err := cast.ToFloat[float32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToFloat32Slice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []float32
		err      bool
	}{
		{nil, nil, true},
		{[]float32{1.5, 2.5}, []float32{1.5, 2.5}, false},
		{[]int{1, 2}, []float32{1.0, 2.0}, false},
		{[]string{"1.5", "2.5"}, []float32{1.5, 2.5}, false},
		{[]string{"1.5", "invalid"}, nil, true},
		{"invalid", nil, true},
		{[]interface{}{math.MaxFloat64}, nil, true}, // Overflow case
	}

	for _, test := range tests {
		result, err := cast.ToFloatSlice[float32](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		err      bool
	}{
		{nil, 0, true},
		{1, 1.0, false},
		{0, 0.0, false},
		{1.5, 1.5, false},
		{"1.5", 1.5, false},
		{"0", 0.0, false},
		{"invalid", 0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{[]int{1}, 0, true},
	}

	for _, test := range tests {
		result, err := cast.ToFloat[float64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToFloat64Slice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []float64
		err      bool
	}{
		{nil, nil, true},
		{[]float64{1.5, 2.5}, []float64{1.5, 2.5}, false},
		{[]int{1, 2}, []float64{1.0, 2.0}, false},
		{[]string{"1.5", "2.5"}, []float64{1.5, 2.5}, false},
		{[]string{"1.5", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToFloatSlice[float64](test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
