package pbio

import (
	"bufio"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

// Encoder encodes protobuf messages and writes to the underlying writer.
type Encoder struct {
	w   io.Writer
	buf []byte
}

// NewEncoder inits a new encoder.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:   w,
		buf: make([]byte, binary.MaxVarintLen64),
	}
}

// Encode encodes a message.
func (e *Encoder) Encode(msg proto.Message) error {
	opt := proto.MarshalOptions{}
	sz := opt.Size(msg)
	e.buf = e.buf[:binary.PutUvarint(e.buf, uint64(sz))]

	data, err := opt.MarshalAppend(e.buf, msg)
	if err != nil {
		return err
	}
	e.buf = data

	_, err = e.w.Write(data)
	return err
}

// --------------------------------------------------------------------

// Decoder decodes protobuf messages from an underlying reader.
type Decoder struct {
	r   *bufio.Reader
	buf []byte
}

// NewDecoder inits a new Decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: bufio.NewReader(r),
	}
}

// Decode decodes a message.
func (d *Decoder) Decode(msg proto.Message) error {
	sz, err := binary.ReadUvarint(d.r)
	if err != nil {
		return err
	}

	if n := int(sz); cap(d.buf) < n {
		d.buf = make([]byte, n)
	} else {
		d.buf = d.buf[:n]
	}
	if _, err := io.ReadFull(d.r, d.buf); err != nil {
		return err
	}

	opt := proto.UnmarshalOptions{}
	return opt.Unmarshal(d.buf, msg)
}
