package goboom

import "encoding/json"

type Error struct {
	StatusCode int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (r Error) Json() []byte {
	jsonString, err := json.MarshalIndent(r, "", "    ")

	if err != nil {
		panic(err)
	}

	if jsonString != nil {
		return jsonString
	}

	return []byte("")
}

func Create(statusCode int, message string, data interface{}) *Error {
	err := &Error{}

	if message != "" {
		err.Message = message
	}

	if statusCode != 0 {
		err.StatusCode = statusCode
	}

	if data != nil {
		err.Payload = data
	}

	return err
}

func InternalServerError(message string, data interface{}) *Error {
	if message == "" {
		message = Error_500
	}
	return Create(500, message, data)
}

func BadRequest(message string, data interface{}) *Error {
	if message == "" {
		message = Error_401
	}
	return Create(400, message, data)
}

func Unauthorized(message string, data interface{}) *Error {
	if message == "" {
		message = Error_401
	}
	return Create(401, message, data)
}

func Wrap(err error, statusCode int, data interface{}) {
	if statusCode == 0 {
		statusCode = 500
	}
	return Create(statusCode, err.Error(), data)
}