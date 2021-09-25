package response

import (
	"webhook/exception"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCError struct {
	err error
}

func (e *GRPCError) Error() string {
	return e.err.Error()
}

func (e *GRPCError) GRPCStatus() *status.Status {
	if err, ok := e.err.(interface {
		Code() exception.ErrorCode
	}); ok {
		return status.New(codeMap[err.Code()], e.Error())
	}
	return status.New(codes.Unknown, e.Error())
}

func Error(err error) error {
	grpcErr := &GRPCError{err}
	status, _ := status.FromError(grpcErr)
	return status.Err()
}

var codeMap = map[exception.ErrorCode]codes.Code{
	exception.ApplicationError: codes.Internal,
}
