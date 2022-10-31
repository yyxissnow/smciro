package xerr

import (
	"github.com/pkg/errors"
)

type XError struct {
	code        int
	userMessage string
	err         error
}

func NewX(code int, msg string) *XError {
	return &XError{
		code:        code,
		userMessage: msg,
		err:         unKnowErr,
	}
}

func New(msg string) error {
	return &XError{
		code: unKnowErrCode,
		err:  errors.New(msg),
	}
}

func Errorf(format string, args ...interface{}) error {
	return &XError{
		code: unKnowErrCode,
		err:  errors.Errorf(format, args),
	}
}

// XWithUMessage 中途记录前端信息
func XWithUMessage(err error, msg string) error {
	if e, ok := err.(*XError); ok {
		e.userMessage = msg
		return e
	}
	return NewX(unKnowErrCode, msg)
}

// XWithError 上层记录前端和堆栈信息
func XWithError(x *XError, err error) *XError {
	if e, ok := err.(*XError); ok && e.userMessage != "" {
		x.userMessage = e.userMessage //保留底册的错误信息
	}
	x.err = err
	return x
}

func XWithMessage(err error, msg string) error {
	return errors.WithMessage(err, msg)
}

func XWithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args)
}

func XWrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func XWrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args)
}

func (x *XError) Error() string {
	if x.err != nil {
		return x.err.Error()
	}
	return x.UMessage()
}

func (x *XError) Code() int {
	return x.code
}

func (x *XError) UMessage() string {
	return x.userMessage
}

func (x *XError) Err() error {
	if x.err == nil {
		return unKnowErr
	}
	return x.err
}

func AnalyseError(err error) *XError {
	if err == nil {
		return Success
	}
	if e, ok := err.(*XError); ok {
		if e.Err() == nil {
			e.err = unKnowErr
		}
		return e
	}
	return errStringToXError(err.Error())
}

func errStringToXError(e string) *XError {
	if e == "" {
		return Success
	}
	return NewX(unKnowErrCode, e)
}
