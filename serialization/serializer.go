package serialization

import (
	"fmt"
	"io"
)

var formatMap = map[Format]FormatSerializer{
	JSON:    &JSONSerializer{},
	MsgPack: &MsgpackSerializer{},
}

// Marshal dumps the struct to bytes in the correct format.
func Marshal(inStruct interface{}, format Format) ([]byte, error) {
	if s, ok := formatMap[format]; ok {
		return s.Marshal(inStruct)
	}
	return nil, fmt.Errorf("unknown format: %s", format)
}

// Unmarshal loads the data in the correct format to the struct.
func Unmarshal(data []byte, outStruct interface{}, format Format) error {
	if s, ok := formatMap[format]; ok {
		return s.Unmarshal(data, outStruct)
	}
	return fmt.Errorf("unknown format: %s", format)
}

// Encode encodes the struct in the correct format & writes it to the writer.
func Encode(inStruct interface{}, w io.Writer, format Format) error {
	if s, ok := formatMap[format]; ok {
		return s.Encode(inStruct, w)
	}
	return fmt.Errorf("unknown format: %s", format)
}

// Decode decodes body from the reader into the struct.
func Decode(r io.Reader, outStruct interface{}, format Format) error {
	if s, ok := formatMap[format]; ok {
		return s.Decode(r, outStruct)
	}
	return fmt.Errorf("unknown format: %s", format)
}
