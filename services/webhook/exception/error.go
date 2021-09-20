package exception

import (
	"fmt"
)

/*
  Error code
*/
const (
	/* Client error codes 001xxx */

	/* Application error codes 002xxx */
	ApplicationError      = "002001"
	UserModelInvalidError = "002002"

	// Database error codes 0021xx
	DatabaseConnectionError = "002101"

	/* Third party API error codes 003xxx */
)

/*
  Error message
*/
var codes = map[string]string{
	ApplicationError:        "Internal server error",
	UserModelInvalidError:   "Faild create or update User model",
	DatabaseConnectionError: "Failed connect database",
}

type Error struct {
	code          string
	message       string
	originalError error
}

func NewError(code string, originalError error) *Error {
	return &Error{code, codes[code], originalError}
}

// output format
func (e *Error) Error() string {
	return fmt.Sprintf("[code: %s] %s %s: %s", e.code, e.outputErrorType(), e.message, e.originalError)
}

func (e *Error) outputErrorType() string {
	switch prefix := e.code[:3]; prefix {
	case "001":
		return "ClientError"
	case "002":
		return "ApplicationError"
	case "003":
		return "DatabaseError"
	default:
		return "UnknownError"
	}
}
