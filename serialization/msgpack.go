package serialization

import (
	"errors"
	"io"

	"github.com/tinylib/msgp/msgp"
)

// MsgpackSerializer implements msgpack serialization
// for the smart serializer.
type MsgpackSerializer struct {
}

// Marshal marshals inStruct to msgpack.
func (m *MsgpackSerializer) Marshal(inStruct interface{}) ([]byte, error) {
	mrsh, ok := inStruct.(msgp.Marshaler)
	if !ok {
		return nil, errors.New("in struct does not implement msgp.Marshaler")
	}
	return mrsh.MarshalMsg(nil)
}

// Unmarshal unmarshals a raw msgpack message to a struct.
func (m *MsgpackSerializer) Unmarshal(rawBytes []byte, outStruct interface{}) error {
	unmarshaler, ok := outStruct.(msgp.Unmarshaler)
	if !ok {
		return errors.New("outStruct does not implement msgp.Unmarshaler")
	}

	_, err := unmarshaler.UnmarshalMsg(rawBytes)
	return err
}

// Encode marshals the struct to a stream.
func (m *MsgpackSerializer) Encode(inStruct interface{}, w io.Writer) error {
	marshaler, ok := inStruct.(msgp.Encodable)
	if !ok {
		return errors.New("inStruct does not implement msgp.Encodable")
	}
	return marshaler.EncodeMsg(msgp.NewWriter(w))
}

// Decode unmarshals the struct from a stream.
func (m *MsgpackSerializer) Decode(r io.Reader, outStruct interface{}) error {
	decoder, ok := outStruct.(msgp.Decodable)
	if !ok {
		return errors.New("outStruct does not implement msgp.Decodable")
	}

	return decoder.DecodeMsg(msgp.NewReader(r))
}
