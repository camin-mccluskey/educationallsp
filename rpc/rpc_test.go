package rpc_test

import (
	"educationallsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	// test case
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	// test case
	incoming := []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}")
	msg, content, err := rpc.DecodeMessage([]byte(incoming))
	contentLength := len(content)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if contentLength != 15 {
		t.Errorf("Expected %d but got %d", 15, contentLength)
	}
	if msg != "hi" {
		t.Errorf("Expected %s but got %s", "Method", msg)
	}
}
