package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestToSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []interface{}
		err      bool
	}{
		{nil, nil, true},
		{[]string{"a", "b"}, []interface{}{"a", "b"}, false},
		{[]int{1, 2}, []any{1, 2}, false},
		{[]float64{1.1, 2.2}, []interface{}{1.1, 2.2}, false},
		{[]bool{true, false}, []interface{}{true, false}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := cast.ToSlice(test.input)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
