package goboom

import (
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create(500, "Internal Server Error", nil)

	if err.StatusCode != 500 {
		t.Error("Expected status code 500 but got", err.StatusCode, "instead")
	}

	if err.Message != "Internal Server Error" {
		t.Error("Expected message 'Internal Server Error' but got", err.Message, "instead")
	}

	if err.Payload != nil {
		t.Error("Expected Payload to be nil but got", err.Payload, "instead")
	}
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError("Something wrong", nil)

	if err.StatusCode != 500 {
		t.Error("Expected status code 500 but got", err.StatusCode, "instead")
	}

	if err.Message != "Something wrong" {
		t.Error("Expected message 'Internal Server Error' but got", err.Message, "instead")
	}

	if err.Payload != nil {
		t.Error("Expected Payload to be nil but got", err.Payload, "instead")
	}
}