package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 别名
type H map[string]interface{}
type Context struct {
	// 请求与回应
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	Params map[string]string
	// 响应信息、状态码
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
func (this *Context)Param(key string)string  {
	value,_:=this.Params[key]
	return value
}
func (this *Context) PostForm(key string) string {
	return this.Req.FormValue(key)
}
func (this *Context) Query(key string) string {
	return this.Req.URL.Query().Get(key)
}
func (this *Context) Status(code int) {
	this.StatusCode = code
	this.Writer.WriteHeader(code)
}
func (this *Context) SetHeader(key string, value string) {
	this.Writer.Header().Set(key, value)
}
func (this *Context) String(code int, format string, values ...interface{}) {
	this.SetHeader("Content-Type", "text/plain")
	this.Status(code)
	this.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
func (this *Context) JSON(code int, obj interface{}) {
	this.SetHeader("Content-Type", "text/plain")
	this.Status(code)
	encoder := json.NewEncoder(this.Writer)     //生成向文件输出的输出流
	if err := encoder.Encode(obj); err != nil { //输出obj
		http.Error(this.Writer, err.Error(), 500)
	}
}
func (this *Context) Data(code int, data []byte) {
	this.Status(code)
	this.Writer.Write(data)
}
func (this *Context) HTML(code int, html string) {
	this.SetHeader("Context-Type", "text/html")
	this.Status(code)
	this.Writer.Write([]byte(html))
}
