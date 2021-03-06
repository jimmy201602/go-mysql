package ast

//定义抽象语法树的基础结构

//所有的节点都有类型信息
//实现所有Node实现的基础信息
type node struct {
	tag    int //节点的类型
	parent ExprNode
}

func Newnode(t int) node {
	return node{tag: t}
}

//实现Set接口
func (n *node) SetTag(tag int) {
	n.tag = tag
}

//实现Set接口
func (n *node) SetParent(e ExprNode) {
	n.parent = e
}

//实现Get接口
func (n *node) GetTag() int {
	return n.tag
}

//实现Get接口
func (n *node) GetParent() ExprNode {
	return n.parent
}
