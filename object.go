package json

import "io"

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

// Get gets string property, returns empty string if property does not exist
func (o *Object) Get(name string) string {
	v, ok := o.data[name]
	if !ok {
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

// TODO more api (GetInt, GetFloat, etc)

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
