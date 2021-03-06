package json

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Value interfaces JSON types
type Value interface {
	json.Marshaler
	// JSON string
	JSON(pretty ...bool) string
	// WriteJSON writes JSON string to given output
	WriteJSON(output io.Writer, pretty ...bool)
}

// Parse returns JSON atom from given input
func Parse(input interface{}) (Value, error) {
	dec, err := NewDecoder(input)
	if err != nil {
		return nil, err
	}

	var val interface{}
	err = dec.Decode(&val)
	if err != nil {
		return nil, err
	}

	switch val.(type) {
	case map[string]interface{}:
		return NewObject(val.(map[string]interface{})), nil
	case []interface{}:
		return NewArray(val.([]interface{})), nil
	default:
		return NewAtom(val), nil
	}
}

// Stringify returns JSON of given value
func Stringify(v interface{}, pretty ...bool) string {
	if len(pretty) > 0 && pretty[0] {
		s, _ := json.MarshalIndent(v, "", "  ")
		return string(s)
	}
	s, _ := json.Marshal(v)
	return string(s)
}

// NewDecoder creates json.Decoder from given input
func NewDecoder(i interface{}) (*json.Decoder, error) {
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
