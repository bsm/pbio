package pbio_test

import (
	"bytes"
	"testing"

	"github.com/bsm/pbio"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

var _ = Describe("Encoder", func() {
	var subject *pbio.Encoder
	var buf *bytes.Buffer

	BeforeEach(func() {
		buf = new(bytes.Buffer)
		subject = pbio.NewEncoder(buf)
	})

	It("should encode", func() {
		Expect(subject.Encode(message)).To(Succeed())
		Expect(buf.Len()).To(BeNumerically("~", 192, 20))
		Expect(subject.Encode(message)).To(Succeed())
		Expect(buf.Len()).To(BeNumerically("~", 384, 20))
	})
})

var _ = Describe("Decoder", func() {
	var subject *pbio.Decoder

	BeforeEach(func() {
		buf := new(bytes.Buffer)
		enc := pbio.NewEncoder(buf)
		Expect(enc.Encode(message)).To(Succeed())
		Expect(enc.Encode(message)).To(Succeed())

		subject = pbio.NewDecoder(buf)
	})

	It("should decode", func() {
		m1 := new(structpb.Value)
		Expect(subject.Decode(m1)).To(Succeed())
		Expect(m1.MarshalJSON()).To(MatchJSON(`{
			"firstName":"John",
			"lastName":"Smith",
			"isAlive":true,
			"age":27,
			"address":{
				"city":"New York",
				"postalCode":"10021-3100",
				"state":"NY",
				"streetAddress":"21 2nd Street"
			}
		}`))

		m2 := new(structpb.Value)
		Expect(subject.Decode(m2)).To(Succeed())

		m3 := new(structpb.Value)
		Expect(subject.Decode(m3)).To(MatchError(`EOF`))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pbio")
}
