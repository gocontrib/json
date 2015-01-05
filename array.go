package json

import "io"

// Array wraps JSON array
type Array struct {
	data []interface{}
}

// NewArray creates new Array instance
func NewArray(arr []interface{}) *Array {
	return &Array{arr}
}

// Get returns value at specified index
func (a *Array) Get(i int) interface{} {
	return a.data[i]
}

// Set value at specified index
func (a *Array) Set(i int, v interface{}) {
	a.data[i] = v
}

// Add new value to array
func (a *Array) Add(item interface{}) {
	a.data = append(a.data, item)
}

// ToArray returns underlying array
func (a *Array) ToArray() []interface{} {
	return a.data
}

// MarshalJSON implements json.Marshaler
func (a *Array) MarshalJSON() ([]byte, error) {
	return []byte(a.JSON()), nil
}

// JSON string
func (a *Array) JSON(pretty ...bool) string {
	return Stringify(&a.data, pretty...)
}

// WriteJSON writes JSON string into given output
func (a *Array) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(a.JSON(pretty...)))
}
