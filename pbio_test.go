package pbio_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math"
	"testing"

	"github.com/bsm/pbio"
	"google.golang.org/protobuf/types/known/structpb"
)

var message, _ = structpb.NewValue(map[string]interface{}{
	"firstName": "John",
	"lastName":  "Smith",
	"isAlive":   true,
	"age":       27,
	"address": map[string]interface{}{
		"streetAddress": "21 2nd Street",
		"city":          "New York",
		"state":         "NY",
		"postalCode":    "10021-3100",
	},
})

func TestEncoder(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := pbio.NewEncoder(buf)

	if err := enc.Encode(message); err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	if exp, got, delta := 192, buf.Len(), 20; math.Abs(float64(exp)-float64(got)) > float64(delta) {
		t.Fatalf("expected %d (±%d), got %d", exp, delta, got)
	}

	if err := enc.Encode(message); err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	if exp, got, delta := 384, buf.Len(), 20; math.Abs(float64(exp)-float64(got)) > float64(delta) {
		t.Fatalf("expected %d (±%d), got %d", exp, delta, got)
	}
}

func TestDecoder(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := pbio.NewEncoder(buf)
	for i := 0; i < 2; i++ {
		if err := enc.Encode(message); err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
	}

	dec := pbio.NewDecoder(buf)
	m1 := new(structpb.Value)
	if err := dec.Decode(m1); err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if data, err := m1.MarshalJSON(); err != nil {
		t.Fatalf("expected no error, but got %v", err)
	} else if exp, got := `{
  "address": {
    "city": "New York",
    "postalCode": "10021-3100",
    "state": "NY",
    "streetAddress": "21 2nd Street"
  },
  "age": 27,
  "firstName": "John",
  "isAlive": true,
  "lastName": "Smith"
}`, normJSON(data); exp != got {
		t.Fatalf("expected:\n%v, but got:\n%v", exp, got)
	}

	m2 := new(structpb.Value)
	if err := dec.Decode(m2); err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	m3 := new(structpb.Value)
	if exp, got := io.EOF, dec.Decode(m3); !errors.Is(got, exp) {
		t.Fatalf("expected %v, but got %v", exp, got)
	}
}

func normJSON(data []byte) string {
	b := new(bytes.Buffer)
	_ = json.Indent(b, data, "", "  ")
	return b.String()
}
