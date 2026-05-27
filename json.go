package orderedmap

import (
	"encoding/json"

	"github.com/mailru/easyjson/jwriter"
)

var (
	_ json.Marshaler   = &OrderedMap[int, any]{}
	_ json.Unmarshaler = &OrderedMap[int, any]{}
)

// MarshalJSON implements the json.Marshaler interface.
func (om *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented" //nolint:funlen
	return nil, nil
}

// this switch takes care of wrapper types around primitive types, such as
// type myType string

// the error is checked at the end of the function

func jsonMarshal(t interface{}, disableHTMLEscape bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode() adds an extra newline, strip it off to guarantee same behavior as json.Marshal

func dumpWriter(writer *jwriter.Writer) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (om *OrderedMap[K, V]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// jsonparser removes the enclosing quotes; we need to restore them to make a valid JSON

// this switch takes care of wrapper types around primitive types, such as
// type myType string

func decodeUTF8(input []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }
