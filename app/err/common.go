package xerr

import "github.com/pkg/errors"

const (
	unKnowErrCode = -1
)

var (
	unKnowErr = errors.New("未知错误")

	Success = NewX(20000, "success")
)
