package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("This code %d already exists. Please try another one", code))
	}
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

var (
	Success                    = NewError(0, "Success")
	ServerError                = NewError(1000000, "Server Error")
	InvalidParams              = NewError(1000001, "Invalid Params")
	NotFound                   = NewError(1000002, "Not Found")
	UnauthorizedAuthNotExist   = NewError(1000003, "UnauthoerizedTokeExists")
	UnauthoerizedTokenError    = NewError(1000004, "Unauthoerized Token Error")
	UnauthoerizedTokenTimeout  = NewError(1000005, "Unauthoerized Token Timeout")
	UnauthoerizedTokenGenerate = NewError(1000006, "Unauthoerized Token Generate")
	TooManyRequests            = NewError(1000007, "Too Many Requests")
	//ServerError                = NewError(1000000, "server error")
)

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d status::%s ", e.Code(), e.Msg())
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg

}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	//for _, d := range details {
	newError.details = append(newError.details, details...)
	//}
	return &newError
}

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
