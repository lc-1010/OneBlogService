package errcode

import (
	pb "github.com/lc-1010/OneBlogService/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).
		WithDetails(&pb.Error{
			Code:    int32(err.Code()),
			Message: err.Msg()})
	return s.Err()
}

func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code

	switch code {
	case ServerError.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case UnauthoerizedTokenError.Code():
		statusCode = codes.Unauthenticated
	case UnauthoerizedTokenTimeout.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case TooManyRequests.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown

	}
	return statusCode
}

type Status struct {
	*status.Status
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(ToRPCCode((code)), msg).
		WithDetails(&pb.Error{Code: int32(code),
			Message: msg})
	return &Status{s}
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}
