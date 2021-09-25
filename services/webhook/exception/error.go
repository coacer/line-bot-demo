package exception

import (
	"errors"
	"fmt"
)

type ErrorCode string

/*
  Error code
*/
const (
	/* Client error codes 001xxx */

	/* Application error codes 002xxx */
	InternalServerError   ErrorCode = "002001"
	UserModelInvalidError ErrorCode = "002002"

	// Database error codes 0021xx
	DatabaseConnectionError ErrorCode = "002101"

	/* Third party API error codes 003xxx */
)

/*
  Error message
*/
var codeMessageMap = map[ErrorCode]string{
	InternalServerError:     "Internal server error",
	UserModelInvalidError:   "Faild create or update User model",
	DatabaseConnectionError: "Failed connect database",
}

type Error struct {
	code          ErrorCode
	message       string
	originalError error
}

func NewError(code ErrorCode, originalError error) *Error {
	if originalError == nil {
		originalError = errors.New("")
	}
	return &Error{code, codeMessageMap[code], originalError}
}

// output format
func (e *Error) Error() string {
	return fmt.Sprintf("[code: %s] %s: %s (%s)", e.code, e.outputErrorType(), e.message, e.originalError)
}

func (e *Error) Code() ErrorCode {
	return e.code
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
