package gee

import (
	"log"
	"net/http"
)

/* 定义函数类型，将函数类型当成一种函数参数
Context结构体组合了请求和响应参数，处理函数改变时，这个定义函数类型也需要做出相应的更改
当由小粒度变成大接口时，也需要做出相应的传参更改 */
type HandlerFunc func(*Context)

// Engine is the uni handler for all requests
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// 键值对保存路由，其他方法调用addRoute方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 实现了接口方法的 struct 都可以强制转换为接口类型。
// 但是手动转换是多余的，因为会自动进行参数转换，所以直接传入engine即可。

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；
// 第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)

}
