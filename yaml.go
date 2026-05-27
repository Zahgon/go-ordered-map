package orderedmap

import (
	"gopkg.in/yaml.v3"
)

var (
	_ yaml.Marshaler   = &OrderedMap[int, any]{}
	_ yaml.Unmarshaler = &OrderedMap[int, any]{}
)

// MarshalYAML implements the yaml.Marshaler interface.
func (om *OrderedMap[K, V]) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// serialize key to yaml, then deserialize it back into the node
// this is a hack to get the correct tag for the key

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (om *OrderedMap[K, V]) UnmarshalYAML(value *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}
