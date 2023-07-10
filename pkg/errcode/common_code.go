package errcode

import (
	"fmt"
	"net/http"
)

// Error repsonse error
type Error struct {
	code    int
	msg     string
	details []string
}

var b_codes = map[int]string{}

// NewError creates a new Error object with the given code and message.
//
// Parameters:
// - code: an integer representing the error code.
// - msg: a string containing the error message.
//
// Returns:
// - *Error: a pointer to the newly created Error object.
func NewError(code int, msg string) *Error {
	if _, ok := b_codes[code]; ok {
		panic(fmt.Sprintf("This code %d already exists. Please try another one", code))
	}
	b_codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

var (
	// Success 表示操作成功
	Success = NewError(0, "Success")

	// ServerError 表示服务器错误
	ServerError = NewError(1000000, "Server Error")

	// InvalidParams 表示参数无效
	InvalidParams = NewError(1000001, "Invalid Params")

	// NotFound 表示资源未找到
	NotFound = NewError(1000002, "Not Found")

	// UnauthorizedAuthNotExist 表示未授权令牌存在
	UnauthorizedAuthNotExist = NewError(1000003, "Unauthorized Token Exists")

	// UnauthoerizedTokenError 表示未授权令牌错误
	UnauthoerizedTokenError = NewError(1000004, "Unauthorized Token Error")

	// UnauthoerizedTokenTimeout 表示未授权令牌超时
	UnauthoerizedTokenTimeout = NewError(1000005, "Unauthorized Token Timeout")

	// UnauthoerizedTokenGenerate 表示生成未授权令牌错误
	UnauthoerizedTokenGenerate = NewError(1000006, "Unauthorized Token Generate")

	// TooManyRequests 表示请求过多
	TooManyRequests  = NewError(1000007, "Too Many Requests")
	MethodNotAllowed = NewError(10000008, "不支持该方法")
	//ServerError                = NewError(1000000, "server error")
)

// Error returns a string representation of the Error object.
//
// No parameters.
// Returns a string.
func (e *Error) Error() string {
	return fmt.Sprintf("code:%d status::%s ", e.Code(), e.Msg())
}

// Code returns the code of the Error.
//
// It has no parameters.
// It returns an integer.
func (e *Error) Code() int {
	return e.code
}

// Msg expects a string representation of the Error object.
func (e *Error) Msg() string {
	return e.msg

}

// Msgf expects a formatted string representation of the Error object.
func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details expects a string representation of the Error object.
func (e *Error) Details() []string {
	return e.details
}

// WithDetails expects a string representation of the Error object.
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	//for _, d := range details {
	newError.details = append(newError.details, details...)
	//}
	return &newError
}

// StatusCode expects a string representation of the Error object.
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthoerizedTokenError.Code():
		fallthrough
	case UnauthoerizedTokenGenerate.Code():
		fallthrough
	case UnauthoerizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
