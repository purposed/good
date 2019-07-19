package serialization

import "io"

// Format represents an available file format.
type Format string

// Available file formats.
const (
	MsgPack Format = "application/msgpack"
	JSON    Format = "application/json"
	YAML    Format = "text/x-yaml"
)

// FormatSerializer defines the method set for a format to be used by the smart serializer.
type FormatSerializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error

	Encode(interface{}, io.Writer) error
	Decode(io.Reader, interface{}) error
}
