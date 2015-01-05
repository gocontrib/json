package json

import . "github.com/onsi/ginkgo"
import . "github.com/onsi/gomega"
import "testing"
import "strings"
import "net/http"

func TestObject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Object")
}

var _ = Describe("Object", func() {
	bob := func(o *Object, err error) {
		Expect(o).NotTo(BeNil())
		Expect(err).To(BeNil())
		Expect(o.NumProperties()).To(Equal(1))
		Expect(o.Get("user")).To(Equal("bob"))
	}

	It("should be parsed from valid JSON string", func() {
		o, err := ParseObject(`{"user": "bob"}`)
		bob(o, err)
	})

	It("should be parsed from reader with valid JSON", func() {
		o, err := ParseObject(strings.NewReader(`{"user": "bob"}`))
		bob(o, err)
	})

	It("should be parsed from http request with valid JSON", func() {
		body := strings.NewReader(`{"user": "bob"}`)
		req, err := http.NewRequest("GET", "/", body)
		o, err := ParseObject(req)
		bob(o, err)
	})

	It("should return error when parse invalid JSON string", func() {
		o, err := ParseObject(`{"user": "bob"`)
		Expect(err).NotTo(BeNil())
		Expect(o).To(BeNil())
	})

	It("JSON string", func() {
		o := NewObject(map[string]interface{}{
			"user": "bob",
		})
		Expect(o.JSON()).To(Equal(`{"user":"bob"}`))
	})

	It("JSON pretty string", func() {
		o := NewObject(map[string]interface{}{
			"user": "bob",
		})
		Expect(o.JSON(true)).To(Equal("{\n  \"user\": \"bob\"\n}"))
	})
})
