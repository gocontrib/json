package json

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Token interfaces JSON types
type Token interface {
	// JSON string
	JSON(pretty ...bool) string
	// WriteJSON writes JSON string to given output
	WriteJSON(output io.Writer, pretty ...bool)
}

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

	dec, err := decoder(input)
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

// TODO get, set API

// JSON returns JSON string
func (o *Object) JSON(pretty ...bool) string {
	if len(pretty) > 0 && pretty[0] {
		s, _ := json.MarshalIndent(&o.data, "", "  ")
		return string(s)
	}
	s, _ := json.Marshal(&o.data)
	return string(s)
}

// WriteJSON writes JSON string to given output
func (o *Object) WriteJSON(w io.Writer, pretty ...bool) {
	w.Write([]byte(o.JSON(pretty...)))
}

func decoder(i interface{}) (*json.Decoder, error) {
	reader, ok := i.(io.Reader)
	if ok {
		return json.NewDecoder(reader), nil
	}

	switch i.(type) {
	case string:
		s := i.(string)
		return json.NewDecoder(strings.NewReader(s)), nil
	case *http.Request:
		req := i.(*http.Request)
		return json.NewDecoder(req.Body), nil
	default:
		return nil, errors.New("invalid input")
	}
}
