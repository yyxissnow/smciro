package ecode

import "strconv"

type sError struct {
	code    int
	message string
}

func NewsError(code int, msg string) *sError {
	return &sError{
		code:    code,
		message: msg,
	}
}

func (s *sError) Error() string {
	return s.message
}

func (s *sError) Code() int {
	return s.code
}

func (s *sError) Message() string {
	return s.message
}

func AnalyseError(err error) *sError {
	if err == nil {
		return Success
	}
	if e, ok := err.(*sError); ok {
		return e
	}
	return errStringTosError(err.Error())
}

func errStringTosError(e string) *sError {
	if e == "" {
		return Success
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		return NewsError(-1, e)
	}
	return NewsError(i, e)
}
