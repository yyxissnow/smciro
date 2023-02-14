package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const (
	LogRecordKey  = "log"
	LogRecordOK   = "ok"
	LogRecordFlag = "flag"
)

type RecordLog struct {
	Formats []LogRecordFormat
	Func    func(*gin.Context, []LogRecordContent)
}

type LogRecordFormat struct {
	ZhFormat string
	EnFormat string
	Flag     string
}

type LogRecordContent struct {
	ZhContent string
	EnContent string
}

func GetLogRecordMap(c *gin.Context) map[string]interface{} {
	value, ok := c.Get(LogRecordKey)
	if !ok {
		return map[string]interface{}{}
	}
	m, ok := value.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	m[LogRecordOK] = LogRecordOK
	return m
}

func LogRecord(c *gin.Context, log RecordLog) {
	value, ok := c.Get(LogRecordKey)
	if !ok {
		return
	}
	lm, ok := value.(map[string]interface{})
	if !ok {
		return
	}
	if lm[LogRecordOK] != LogRecordOK {
		return
	}
	delete(lm, LogRecordOK)
	var flag = lm[LogRecordFlag]
	delete(lm, LogRecordFlag)
	var contents []LogRecordContent

	for _, format := range log.Formats {
		if format.Flag != flag {
			continue
		}
		zh := format.ZhFormat
		en := format.EnFormat
		for name, value := range lm {
			zh = strings.Replace(zh, "{"+name+"}", cast.ToString(value), 1)
			en = strings.Replace(en, "{"+name+"}", cast.ToString(value), 1)
		}
		contents = append(contents, LogRecordContent{zh, en})
	}

	log.Func(c, contents)
}
