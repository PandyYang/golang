package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

//接收者 实现是node.print()
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

//普通方式 实现是print(node)
func print2(node treeNode) {
	fmt.Print(node.value)
}

//go是值传递 使用指针作为方法的接收者
func (node *treeNode) setValue(value int){
	node.value = value
}

//使用工厂函数 实现构造函数功能
	//返回了局部变量地址 在C++中是典型错误 但是在go语言中可以
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

func (node *treeNode) traverse(){
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main(){
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	//无论是地址还是指针 使用.进行成员的访问
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.print()

	root.left.right.setValue(6)
	root.left.setValue(6)

	fmt.Println()

	root.print()
	root.setValue(100)

	pRoot := &root
	pRoot.print()
	pRoot.setValue(300)
	pRoot.print()

	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)

	root.traverse()
	}
