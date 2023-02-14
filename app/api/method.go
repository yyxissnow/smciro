package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R *gin.RouterGroup
}

type Method struct {
	Middlewares []gin.HandlerFunc
	RecordLog   RecordLog
}

type Option interface {
	apply(*Method)
}

type optionFunc func(*Method)

func (f optionFunc) apply(m *Method) {
	f(m)
}

func WithMiddleware(middlewares ...gin.HandlerFunc) Option {
	return optionFunc(func(method *Method) {
		method.Middlewares = append(method.Middlewares, middlewares...)
	})
}

func WithLogRecord(log RecordLog) Option {
	return optionFunc(func(method *Method) {
		method.RecordLog = log
	})
}

func withOptions(ops ...Option) *Method {
	method := new(Method)
	for _, o := range ops {
		o.apply(method)
	}
	return method
}

func (a *Handler) Group(path string, middlewares ...gin.HandlerFunc) *Handler {
	return &Handler{a.R.Group(path, middlewares...)}
}

func (a *Handler) Use(middlewares ...gin.HandlerFunc) {
	a.R.Use(middlewares...)
}

func (a *Handler) POST(path string, handle gin.HandlerFunc, ops ...Option) {
	method := withOptions(ops...)
	method.Middlewares = append(method.Middlewares, func(c *gin.Context) { loader(c, handle, method.RecordLog) })
	a.R.POST(path, method.Middlewares...)
}

func (a *Handler) GET(path string, handle gin.HandlerFunc, ops ...Option) {
	method := withOptions(ops...)
	method.Middlewares = append(method.Middlewares, func(c *gin.Context) { loader(c, handle, method.RecordLog) })
	a.R.GET(path, method.Middlewares...)
}

func (a *Handler) PUT(path string, handle gin.HandlerFunc, ops ...Option) {
	method := withOptions(ops...)
	method.Middlewares = append(method.Middlewares, func(c *gin.Context) { loader(c, handle, method.RecordLog) })
	a.R.PUT(path, method.Middlewares...)
}

func (a *Handler) DELETE(path string, handle gin.HandlerFunc, ops ...Option) {
	method := withOptions(ops...)
	method.Middlewares = append(method.Middlewares, func(c *gin.Context) { loader(c, handle, method.RecordLog) })
	a.R.DELETE(path, method.Middlewares...)
}

func (a *Handler) Any(path string, handle gin.HandlerFunc, ops ...Option) {
	method := withOptions(ops...)
	method.Middlewares = append(method.Middlewares, func(c *gin.Context) { loader(c, handle, method.RecordLog) })
	a.R.Any(path, method.Middlewares...)
}

func (a *Handler) StaticFS(relativePath string, fs http.FileSystem) {
	a.R.StaticFS(relativePath, fs)
}

func (a *Handler) Static(relativePath, root string) {
	a.R.Static(relativePath, root)
}

func loader(c *gin.Context, handle gin.HandlerFunc, log RecordLog) {
	c.Set(LogRecordKey, map[string]interface{}{})
	handle(c)
	LogRecord(c, log)
}
