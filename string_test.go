package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		err      bool
	}{
		{nil, "", true},
		{"hello", "hello", false},
		{123, "123", false},
		{123.45, "123.45", false},
		{true, "true", false},
		{false, "false", false},
		{[]int{1, 2, 3}, "", true},
	}

	for _, test := range tests {
		result, err := cast.ToString(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []string
		err      bool
	}{
		{nil, nil, true},
		{[]string{"a", "b"}, []string{"a", "b"}, false},
		{[]int{1, 2}, []string{"1", "2"}, false},
		{[]float64{1.1, 2.2}, []string{"1.1", "2.2"}, false},
		{[]bool{true, false}, []string{"true", "false"}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToStringSlice(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
