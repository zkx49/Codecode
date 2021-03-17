// 设计tire树，匹配动态路由
package gee
// 树的结点,重在理解各字段的含义
import (
	"fmt"
	"strings"
)
type node struct{
	pattern string    //待匹配路由？
	part string       //路由中的一部分
	children []*node  //子节点
	isWild bool       //是否精确匹配，part含有：或* 时为true
}
// %t,以true或false的方式输出布尔值
// %T  输出值的类型
func (this *node)String() string {
	return fmt.Sprintf("node{pattern=%s,part=%s,isWild=%t}",this.pattern,this.part,this.isWild)
}

// 插入结点   递归插入  重点在于算法跟实际应用的转换
func (this *node)insert(pattern string,parts []string,height int)  {
	if len(parts)==height {
		this.pattern=pattern
		return
	}
	part:=parts[height]
	child:=this.matchChild(part)
	if child==nil {
		child=&node{part:part,isWild:part[0]==':'||part[0]=='*'}
		this.children=append(this.children,child)
	}
	child.insert(pattern,parts,height+1)
}
// 递归查找
func (this *node)search(parts []string,height int) *node {
	if len(parts)==height || strings.HasPrefix(this.part,"*") {
		if this.pattern == "" {
			return nil
		}
		return this
	}
	part:=parts[height]
	children:=this.matchChildren(part)
	for _, child := range children {
		result:=child.search(parts,height+1)
		if result !=nil {
			return result
		}
	}
	return nil
}
func (this *node)travel(list *([]*node))  {
	if this.pattern != "" {
		*list=append(*list,this)
	}
	for _, child := range this.children {
		child.travel(list)
	}
}

// 两个匹配方法
// 匹配子节点中路由部分
func (this *node)matchChild(part string)  *node{
	for _, child := range this.children {
		if child.part==part || child.isWild{
			return child
		}
	}
	return nil
}
// 查找所有匹配成功的结点
func (this *node)matchChildren(part string)  []*node{
	nodes:=make([]*node, 0)
	for _, child := range this.children {
		if child.part==part || child.isWild {
			nodes=append(nodes,child)
		}
	}
	return nodes
}