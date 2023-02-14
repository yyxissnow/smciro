package xerr

import (
	"github.com/pkg/errors"
)

type XError struct {
	code    int
	message string
	err     error
}

func NewX(code int, msg string) *XError {
	return &XError{
		code:    code,
		message: msg,
		err:     unKnowErr,
	}
}

func New(msg string) *XError {
	return &XError{
		code: unKnowErrCode,
		err:  errors.New(msg),
	}
}

func Errorf(format string, args ...interface{}) *XError {
	return &XError{
		code: unKnowErrCode,
		err:  errors.Errorf(format, args),
	}
}

// WithError 底层错误码透传到上层
func (x *XError) WithError(err error) {
	if e, ok := err.(*XError); ok && e.code != unKnowErrCode {
		x.code = e.code
		x.message = e.message
	}
	x.err = err
}

func WithMessage(err error, msg string) {
	if x, ok := err.(*XError); ok {
		x.err = errors.WithMessage(x.err, msg)
	}
}

func XWithMessagef(err error, format string, args ...interface{}) {
	if x, ok := err.(*XError); ok {
		x.err = errors.WithMessagef(x.err, format, args...)
	}
}

func Wrap(err error, message string) *XError {
	return &XError{
		code: unKnowErrCode,
		err:  errors.Wrap(err, message),
	}
}

func Wrapf(err error, format string, args ...interface{}) *XError {
	return &XError{
		code: unKnowErrCode,
		err:  errors.Wrapf(err, format, args),
	}
}

func (x *XError) Error() string {
	if x.err != nil {
		return x.err.Error()
	}
	return x.Message()
}

func (x *XError) Code() int {
	return x.code
}

func (x *XError) Message() string {
	return x.message
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

func errStringToXError(msg string) *XError {
	if msg == "" {
		return Success
	}
	return NewX(unKnowErrCode, msg)
}
