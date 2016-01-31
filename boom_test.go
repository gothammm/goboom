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

func TestBadRequest(t *testing.T) {
	err := BadRequest("Bad Request m8", nil)

	if err.StatusCode != 400 {
		t.Error("Expected status code 400 but got", err.StatusCode, "instead")
	}

	if err.Message != "Bad Request m8" {
		t.Error("Expected message 'Bad Request m8' but got", err.Message, "instead")
	}

	if err.Payload != nil {
		t.Error("Expected Payload to be nil but got", err.Payload, "instead")
	}
}

func TestUnauthorized(t *testing.T) {
	err := Unauthorized("You shall not pass", nil)

	if err.StatusCode != 401 {
		t.Error("Expected status code 401 but got", err.StatusCode, "instead")
	}

	if err.Message != "You shall not pass" {
		t.Error("Expected message 'You shall not pass' but got", err.Message, "instead")
	}

	if err.Payload != nil {
		t.Error("Expected Payload to be nil but got", err.Payload, "instead")
	}
}