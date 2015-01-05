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

// Atom is primitive JSON value (number, string, bool, null)
type Atom struct {
	data interface{}
}

// NewAtom creates new Atom instance
func NewAtom(v interface{}) *Atom {
	return &Atom{v}
}

// Unwrap underlying value
func (a *Atom) Unwrap() interface{} {
	return a.data
}

// String coerces into a string
func (a *Atom) String() (string, error) {
	if s, ok := (a.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Int coerces into an int
func (a *Atom) Int() (int, error) {
	return toInt(a.data)
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
	return 0, errors.New("invalid Atom type")
}

// Int64 coerces into an int64
func (a *Atom) Int64() (int64, error) {
	return toInt64(a.data)
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
func (a *Atom) Uint64() (uint64, error) {
	return toUint64(a.data)
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
func (a *Atom) Float64() (float64, error) {
	return toFloat64(a.data)
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
func (a *Atom) JSON(pretty ...bool) string {
	return Stringify(&a.data, pretty...)
}

// WriteJSON writes JSON string into given output
func (a *Atom) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(a.JSON(pretty...)))
}

// MarshalJSON implements json.Marshaler
func (a *Atom) MarshalJSON() ([]byte, error) {
	return []byte(a.JSON()), nil
}
