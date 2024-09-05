package error

import (
	"fmt"
)

var (
	ErrServerInternal = &Error{
		Code:       20001,
		HttpStatus: 500,
		Msg:        "Server Internal Error",
	}
	ErrBadRequest = &Error{
		Code:       20002,
		HttpStatus: 400,
		Msg:        "Bad Request",
	}
	ErrUnauthorized = &Error{
		Code:       20003,
		HttpStatus: 401,
		Msg:        "Unauthorized",
	}
	ErrorNotFound = &Error{
		Code:       20004,
		HttpStatus: 404,
		Msg:        "Not Found",
	}
	ErrForbidden = &Error{
		Code:       20005,
		HttpStatus: 403,
		Msg:        "Forbidden",
	}

	ErrRegionNotFound = &Error{
		Code:       30005,
		HttpStatus: 403,
		Msg:        "Region not found",
	}

	ErrDestinationNotFound = &Error{
		Code:       40005,
		HttpStatus: 403,
		Msg:        "Region not found",
	}
)

type Error struct {
	HttpStatus  int    `json:"-" swaggerignore:"true"`
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	InternalMsg string `json:"-" swaggerignore:"true"`
}

func (e *Error) StatusCode() int {
	return e.HttpStatus
}

func (e *Error) Wrap(err error) *Error {
	res := &Error{
		HttpStatus: e.HttpStatus,
		Code:       e.Code,
		Msg:        e.Msg,
	}
	if res.Code == ErrServerInternal.Code {
		res.InternalMsg = fmt.Sprintf("%s", err.Error())
		return res
	}
	res.Msg = fmt.Sprintf("%s: %s", e.Msg, err.Error())
	return res
}

func (e *Error) Format(errorMessage string, params ...any) *Error {
	errorMessage = fmt.Sprintf(errorMessage, params...)
	res := &Error{
		HttpStatus: e.HttpStatus,
		Code:       e.Code,
		Msg:        e.Msg,
	}
	if res.Code == ErrServerInternal.Code {
		res.InternalMsg = fmt.Sprintf("%s", errorMessage)
		return res
	}
	res.Msg = fmt.Sprintf("%s: %s", e.Msg, errorMessage)
	return res
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
	if e.InternalMsg != "" {
		msg = fmt.Sprintf("%s, internal msg: %s", msg, e.InternalMsg)
	}
	return msg
}
