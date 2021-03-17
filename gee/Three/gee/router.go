package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
// 解析待匹配部分
func parsePattern(pattern string)[]string {
	ss:=strings.Split(pattern,"/")
	parts:=make([]string, 0)
	for _, item := range ss {
		if item!="" {
			parts=append(parts,item)
			if item[0] == '*'{
				break   //跳出循环
			}
		}
	}
	return parts
}
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts:=parsePattern(pattern)
	key := method + "-" + pattern
	_,ok:=r.roots[method]
	if !ok {
		r.roots[method]=&node{}
	}
	r.roots[method].insert(pattern,parts,0)
	r.handlers[key] = handler
}

func (r *router)getRoute(method string,path string,) (*node,map[string]string) {
	searchparts:=parsePattern(path)
	params:=make(map[string]string)
	root,ok:=r.roots[method]   //查找方法，放回结点
	if !ok {
		return nil,nil
	}
	n:=root.search(searchparts,0)
	if n!=nil {
		parts:=parsePattern(n.pattern)
		for index,part:=range parts{
			if part[0]==':' {
				params[part[1:]]=searchparts[index]
			}
			if part[0]=='*'&& len(part)>1 {
				params[part[1:]]=strings.Join(searchparts[index:],"/")
				break
			}
		}
		return n,params
	}
	return nil,nil
}
func (r *router)getRoutes(method string)[]*node  {
	root,ok:=r.roots[method]
	if !ok {
		return nil
	}
	nodes:=make([]*node, 0)
	root.travel(&nodes)
	return nodes
}
func (r *router) handle(c *Context) {
	n,params:=r.getRoute(c.Method,c.Path)
	if n!=nil {
		c.Params=params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
	
}
