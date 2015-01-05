package json

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
)

var (
	errInvalidType = errors.New("invalid value type")
)

// Value is primitive JSON value (number, string, bool, null)
type Value struct {
	data interface{}
}

// NewValue creates new value
func NewValue(v interface{}) *Value {
	return &Value{v}
}

// Unwrap underlying value
func (v *Value) Unwrap() interface{} {
	return v.data
}

// String coerces into a string
func (v *Value) String() (string, error) {
	if s, ok := (v.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Int coerces into an int
func (v *Value) Int() (int, error) {
	return toInt(v.data)
}

// coerces into an int
func toInt(i interface{}) (int, error) {
	switch i.(type) {
	case json.Number:
		i, err := i.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(i).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(i).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(i).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int64 coerces into an int64
func (v *Value) Int64() (int64, error) {
	return toInt64(v.data)
}

func toInt64(i interface{}) (int64, error) {
	switch i.(type) {
	case json.Number:
		return i.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(i).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(i).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(i).Uint()), nil
	}
	return 0, errInvalidType
}

// Uint64 coerces into an uint64
func (v *Value) Uint64() (uint64, error) {
	return toUint64(v.data)
}

func toUint64(i interface{}) (uint64, error) {
	switch i.(type) {
	case json.Number:
		return strconv.ParseUint(i.(json.Number).String(), 10, 64)
	case float32, float64:
		return uint64(reflect.ValueOf(i).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(i).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(i).Uint(), nil
	}
	return 0, errInvalidType
}

// Float64 coerces into a float64
func (v *Value) Float64() (float64, error) {
	return toFloat64(v.data)
}

func toFloat64(i interface{}) (float64, error) {
	switch i.(type) {
	case json.Number:
		return i.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(i).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(i).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(i).Uint()), nil
	}
	return 0, errInvalidType
}

// JSON string
func (v *Value) JSON(pretty ...bool) string {
	return Stringify(&v.data, pretty...)
}

// WriteJSON writes JSON string into given output
func (v *Value) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(v.JSON(pretty...)))
}

// MarshalJSON implements json.Marshaler
func (v *Value) MarshalJSON() ([]byte, error) {
	return []byte(v.JSON()), nil
}
