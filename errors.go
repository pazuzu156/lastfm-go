package lastfm

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// APIResponseStatusOk represents an OK status.
	APIResponseStatusOk = "ok"

	// APIResponseStatusFailed represents a FAILED status.
	APIResponseStatusFailed = "failed"
)

const (
	// ErrorAuthRequired represents an auth required error code.
	ErrorAuthRequired = 50

	// ErrorParameterMissing represents a missing parameter error code.
	ErrorParameterMissing = 51

	// ErrorInvalidTypeOfArgument represents an invalid agument type error code.
	ErrorInvalidTypeOfArgument = 52
)

// Messages contains API error messages
var Messages = map[int]string{
	50: "This method requires authentication.",
	51: "Required parameter missing. Required: %v, Missing: %v.",
	52: "Invalid type of argument passed. Supported types are int, string and []string.",
}

// Error is an API error struct.
type Error struct {
	Code    int
	Message string
	Where   string
	Caller  string
}

// Error returns an error string with code.
func (e *Error) Error() string {
	return fmt.Sprintf("LastfmError[%d]: %s (%s)", e.Code, e.Message, e.Caller)
}

func newAPIError(errorXML *APIError) (e *Error) {
	e = new(Error)
	e.Code = errorXML.Code
	e.Message = strings.TrimSpace(errorXML.Message)
	return e
}

func newLibError(code int, message string) (e *Error) {
	e = new(Error)
	e.Code = code
	e.Message = message
	return e
}

func appendCaller(err error, caller string) {
	if err != nil && reflect.TypeOf(err).String() == "*lastfm.LastfmError" {
		err.(*Error).Caller = caller
	}
}
