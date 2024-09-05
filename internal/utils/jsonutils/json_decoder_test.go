package jsonutils

import (
	"log"
	"testing"
)

var (
	testJsonString = "{\"name\":\"bar\", \"object\": {\"name\": \"foo\"}},{\"name\":\"foo\", \"object\": {\"name\": \"bar\"}}"
)

func TestJsonStreamDecoder(t *testing.T) {
	dec := NewJsonStreamDecoder()

	result1 := dec.Write(testJsonString[:50])
	log.Println(result1)

	result2 := dec.Write(testJsonString[50:])
	log.Println(result2)

}
