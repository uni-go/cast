package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		err      bool
	}{
		{nil, false, true},
		{true, true, false},
		{false, false, false},
		{1, true, false},
		{0, false, false},
		{1.5, true, false},
		{0.0, false, false},
		{"true", true, false},
		{"false", false, false},
		{"invalid", false, true},
		{[]int{1}, false, true},
	}

	for _, test := range tests {
		result, err := cast.ToBool(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToBoolSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []bool
		err      bool
	}{
		{nil, nil, true},
		{[]bool{true, false}, []bool{true, false}, false},
		{[]int{1, 0}, []bool{true, false}, false},
		{[]string{"true", "false"}, []bool{true, false}, false},
		{[]string{"true", "invalid"}, nil, true},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToBoolSlice(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
