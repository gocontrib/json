package json

import (
	"errors"
	"io"
)

var (
	errNoProperty = errors.New("property not exist")
)

// Object is generic JSON object
type Object struct {
	data map[string]interface{}
}

// NewObject creates new JObject
func NewObject(data map[string]interface{}) *Object {
	if data == nil {
		data = make(map[string]interface{})
	}
	return &Object{
		data: data,
	}
}

// ParseObject parses JSON object from given input
func ParseObject(input interface{}) (*Object, error) {
	obj := NewObject(nil)

	dec, err := NewDecoder(input)
	if err != nil {
		return nil, err
	}

	err = dec.Decode(&obj.data)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// NumProperties returns number of properties.
func (o *Object) NumProperties() int {
	return len(o.data)
}

// Get property value
func (o *Object) Get(name string) interface{} {
	v, ok := o.data[name]
	if !ok {
		return nil
	}

	a, ok := v.(Atom)
	if ok {
		return a.data
	}

	return v
}

// Gets string property, returns empty string if property does not exist
func (o *Object) Gets(name string) string {
	v := o.Get(name)
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

// Set sets given property
func (o *Object) Set(name string, val interface{}) {
	o.data[name] = val
}

// Del removes specified property
func (o *Object) Del(name string) {
	delete(o.data, name)
}

// Bool gets bool property.
func (o *Object) Bool(name string) (bool, error) {
	v := o.Get(name)
	if s, ok := v.(bool); ok {
		return s, nil
	}
	return false, errors.New("type assertion to bool failed")
}

// Int gets int property
func (o *Object) Int(name string) (int, error) {
	v := o.Get(name)
	if v == nil {
		return 0, errNoProperty
	}
	return toInt(v)
}

// Int64 gets int64 property
func (o *Object) Int64(name string) (int64, error) {
	v := o.Get(name)
	if v == nil {
		return 0, errNoProperty
	}
	return toInt64(v)
}

// Uint64 gets uint64 property
func (o *Object) Uint64(name string) (uint64, error) {
	v := o.Get(name)
	if v == nil {
		return 0, errNoProperty
	}
	return toUint64(v)
}

// Float64 gets float64 property
func (o *Object) Float64(name string) (float64, error) {
	v := o.Get(name)
	if v == nil {
		return 0, errNoProperty
	}
	return toFloat64(v)
}

// JSON returns JSON string
func (o *Object) JSON(pretty ...bool) string {
	return Stringify(&o.data, pretty...)
}

// WriteJSON writes JSON string to given output
func (o *Object) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(o.JSON(pretty...)))
}

// MarshalJSON implements json.Marshaler
func (o *Object) MarshalJSON() ([]byte, error) {
	return []byte(o.JSON()), nil
}
