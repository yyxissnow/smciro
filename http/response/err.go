package response

import "fmt"

type ErrorCode int

var errMap map[ErrorCode]string
var successCode int

func init() {
	errMap = make(map[ErrorCode]string)
}

func LoadErrMap(success int, m map[ErrorCode]string) {
	for code, msg := range m {
		if _, ok := errMap[code]; ok {
			panic(fmt.Sprintf("const code %v already exist", code))
		}
		errMap[code] = msg
	}
	successCode = success
}

func getErrMessage(code ErrorCode) string {
	if msg, ok := errMap[code]; ok {
		return msg
	}
	return "code msg not find"
}
