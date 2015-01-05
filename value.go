package json

import "io"

// Value is primitive JSON value (number, string, bool, null)
type Value struct {
	val interface{}
}

// NewValue creates new value
func NewValue(v interface{}) *Value {
	return &Value{v}
}

// TODO more api (String, MustString, Int, MustInt, etc)

// JSON string
func (v *Value) JSON(pretty ...bool) string {
	return Stringify(&v.val, pretty...)
}

// WriteJSON writes JSON string into given output
func (v *Value) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(v.JSON(pretty...)))
}

// MarshalJSON implements json.Marshaler
func (v *Value) MarshalJSON() ([]byte, error) {
	return []byte(v.JSON()), nil
}
