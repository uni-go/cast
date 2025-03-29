# Cast Package Documentation

The `cast` package provides utilities for type casting and conversion in Go. It includes functions for converting values to various Go types, such as `bool`, `int`, `uint`, `float`, `string`, and their respective slices. Additionally, it offers a `Caster` interface for more structured type conversion.

## Installation

To install the `cast` package, use the following command:

```sh
go get github.com/uni-go/cast
```

Then, import the package in your Go code:

```go
import "github.com/uni-go/cast"
```

## Functions

### `ToBool`

Converts an interface to a `bool`.  
**Signature**:

```go
func ToBool(value interface{}) (bool, error)
```

**Example**:

```go
result, err := cast.ToBool("true")
fmt.Println(result) // Output: true
```

### `ToBoolSlice`

Converts an interface to a slice of `bool`.  
**Signature**:

```go
func ToBoolSlice(value interface{}) ([]bool, error)
```

**Example**:

```go
result, err := cast.ToBoolSlice([]string{"true", "false"})
fmt.Println(result) // Output: [true false]
```

### `ToSigned`

Converts an interface to a signed integer type (`int`, `int8`, `int16`, `int32`, `int64`).  
**Signature**:

```go
func ToSigned[T int | int8 | int16 | int32 | int64](value interface{}) (T, error)
```

**Example**:

```go
result, err := cast.ToSigned[int]("123")
fmt.Println(result) // Output: 123
```

### `ToSignedSlice`

Converts an interface to a slice of signed integers.  
**Signature**:

```go
func ToSignedSlice[T int | int8 | int16 | int32 | int64](value interface{}) ([]T, error)
```

**Example**:

```go
result, err := cast.ToSignedSlice[int]([]string{"1", "2"})
fmt.Println(result) // Output: [1 2]
```

### `ToUnsigned`

Converts an interface to an unsigned integer type (`uint`, `uint8`, `uint16`, `uint32`, `uint64`).  
**Signature**:

```go
func ToUnsigned[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) (T, error)
```

**Example**:

```go
result, err := cast.ToUnsigned[uint]("123")
fmt.Println(result) // Output: 123
```

### `ToUnsignedSlice`

Converts an interface to a slice of unsigned integers.  
**Signature**:

```go
func ToUnsignedSlice[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) ([]T, error)
```

**Example**:

```go
result, err := cast.ToUnsignedSlice[uint]([]string{"1", "2"})
fmt.Println(result) // Output: [1 2]
```

### `ToFloat`

Converts an interface to a float type (`float32`, `float64`).  
**Signature**:

```go
func ToFloat[T float32 | float64](value interface{}) (T, error)
```

**Example**:

```go
result, err := cast.ToFloat[float64]("123.45")
fmt.Println(result) // Output: 123.45
```

### `ToFloatSlice`

Converts an interface to a slice of floats.  
**Signature**:

```go
func ToFloatSlice[T float32 | float64](value interface{}) ([]T, error)
```

**Example**:

```go
result, err := cast.ToFloatSlice[float64]([]string{"1.1", "2.2"})
fmt.Println(result) // Output: [1.1 2.2]
```

### `ToString`

Converts an interface to a `string`.  
**Signature**:

```go
func ToString(value interface{}) (string, error)
```

**Example**:

```go
result, err := cast.ToString(123)
fmt.Println(result) // Output: "123"
```

### `ToStringSlice`

Converts an interface to a slice of `string`.  
**Signature**:

```go
func ToStringSlice(value interface{}) ([]string, error)
```

**Example**:

```go
result, err := cast.ToStringSlice([]int{1, 2})
fmt.Println(result) // Output: ["1" "2"]
```

### `ToSlice`

Converts an interface to a slice of `interface{}`.  
**Signature**:

```go
func ToSlice(value interface{}) ([]interface{}, error)
```

**Example**:

```go
result, err := cast.ToSlice([]int{1, 2})
fmt.Println(result) // Output: [1 2]
```

## Caster Interface

The `Caster` interface provides methods for type casting and conversion. It allows structured and reusable type conversions with fallback mechanisms.

### Methods

- **`IsNil() bool`**: Checks if the value is nil.
- **`Interface() interface{}`**: Returns the value as an `interface{}`.
- **`Bool() (bool, error)`**: Converts the value to a `bool`.
- **`BoolSafe(fallback bool) bool`**: Converts the value to a `bool`, returning a fallback value on error.
- **`BoolSlice() ([]bool, error)`**: Converts the value to a slice of `bool`.
- **`BoolSliceSafe(fallback []bool) []bool`**: Converts the value to a slice of `bool`, returning a fallback value on error.
- **`Int() (int, error)`**: Converts the value to an `int`.
- **`IntSafe(fallback int) int`**: Converts the value to an `int`, returning a fallback value on error.
- **`IntSlice() ([]int, error)`**: Converts the value to a slice of `int`.
- **`IntSliceSafe(fallback []int) []int`**: Converts the value to a slice of `int`, returning a fallback value on error.
- **`Float64() (float64, error)`**: Converts the value to a `float64`.
- **`Float64Safe(fallback float64) float64`**: Converts the value to a `float64`, returning a fallback value on error.
- **`String() (string, error)`**: Converts the value to a `string`.
- **`StringSafe(fallback string) string`**: Converts the value to a `string`, returning a fallback value on error.

### Example Usage

```go
caster := cast.NewCaster("123")

// Convert to int
intValue, err := caster.Int()
fmt.Println(intValue) // Output: 123

// Convert to bool with fallback
boolValue := caster.BoolSafe(false)
fmt.Println(boolValue) // Output: true

// Convert to string slice
stringSlice, err := caster.StringSlice()
fmt.Println(stringSlice) // Output: ["1" "2" "3"]
```

## Error Handling

The package provides utility functions to identify specific error types:

- **`IsNilError(err error) bool`**: Checks if the error is due to a nil value.
- **`IsCastError(err error) bool`**: Checks if the error is due to an invalid type conversion.
- **`IsOverflowError(err error) bool`**: Checks if the error is due to a value overflow.

---

This documentation provides a comprehensive overview of the `cast` package and its capabilities. For more details, refer to the source code or examples provided in the repository.
