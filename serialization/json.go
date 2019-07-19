package serialization

import (
	"encoding/json"
	"io"
)

// JSONSerializer serializes messages to json.
type JSONSerializer struct{}

// Marshal marshals inStruct to msgpack.
func (m *JSONSerializer) Marshal(inStruct interface{}) ([]byte, error) {
	return json.Marshal(inStruct)
}

// Unmarshal unmarshals a raw msgpack message to a struct.
func (m *JSONSerializer) Unmarshal(rawBytes []byte, outStruct interface{}) error {
	return json.Unmarshal(rawBytes, outStruct)
}

// Encode marshals the struct to a stream.
func (m *JSONSerializer) Encode(inStruct interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(inStruct)
}

// Decode unmarshals the struct from a stream.
func (m *JSONSerializer) Decode(r io.Reader, outStruct interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(outStruct)
}
