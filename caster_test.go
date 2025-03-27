package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uni-go/cast"
)

func TestIsNil(t *testing.T) {
	c := cast.NewCaster(nil)
	assert.Equal(t, true, c.IsNil())
}

func TestInterface(t *testing.T) {
	value := 42
	c := cast.NewCaster(value)
	assert.Equal(t, value, c.Interface())
}

func TestSlice(t *testing.T) {
	value := []interface{}{1, "hi", 3.1415}
	c := cast.NewCaster(value)
	res, err := c.Slice()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestSliceSafe(t *testing.T) {
	value := "not a slice"
	fallback := []interface{}{1, 2, 3}
	c := cast.NewCaster(value)
	res := c.SliceSafe(fallback)
	assert.Equal(t, fallback, res)
}

func TestUnmarshal(t *testing.T) {
	jsonData := `{"key": "value"}`
	c := cast.NewCaster(jsonData)
	var out map[string]string
	err := c.Unmarshal(&out)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"key": "value"}, out)
}

func TestBool(t *testing.T) {
	value := true
	c := cast.NewCaster(value)
	res, err := c.Bool()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestBoolSafe(t *testing.T) {
	value := "not a bool"
	fallback := true
	c := cast.NewCaster(value)
	res := c.BoolSafe(fallback)
	assert.Equal(t, fallback, res)
}

func TestInt(t *testing.T) {
	value := 42
	c := cast.NewCaster(value)
	res, err := c.Int()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestIntSafe(t *testing.T) {
	value := "not an int"
	fallback := 42
	c := cast.NewCaster(value)
	res := c.IntSafe(fallback)
	assert.Equal(t, fallback, res)
}

func TestFloat64(t *testing.T) {
	value := 3.1415
	c := cast.NewCaster(value)
	res, err := c.Float64()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestFloat64Safe(t *testing.T) {
	value := "not a float"
	fallback := 3.1415
	c := cast.NewCaster(value)
	res := c.Float64Safe(fallback)
	assert.Equal(t, fallback, res)
}

func TestString(t *testing.T) {
	value := "hello"
	c := cast.NewCaster(value)
	res, err := c.String()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestStringSafe(t *testing.T) {
	value := 123
	fallback := "fallback"
	c := cast.NewCaster(value)
	res := c.StringSafe(fallback)
	assert.Equal(t, "123", res)
}

func TestStringSlice(t *testing.T) {
	value := []string{"a", "b", "c"}
	c := cast.NewCaster(value)
	res, err := c.StringSlice()
	assert.NoError(t, err)
	assert.Equal(t, value, res)
}

func TestStringSliceSafe(t *testing.T) {
	value := "not a slice"
	fallback := []string{"fallback"}
	c := cast.NewCaster(value)
	res := c.StringSliceSafe(fallback)
	assert.Equal(t, fallback, res)
}
