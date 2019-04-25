package treeMapreduce

import "fmt"

type Data struct {
	Flag int
	N string
	A string
	T string
	Q string
}

type Node struct {
	Data
	Left *Node
	Right * Node
}

func (node *Node) MiddleOrder() {
	if node == nil {
		return
	}
	node.Left.MiddleOrder()
	node.Print()
	node.Right.MiddleOrder()
}

func (node *Node) Print(){
	fmt.Printf("姓名：%s\n", node.Data.N)
	fmt.Printf("地区：%s\n", node.Data.A)
	fmt.Printf("电话：%s\n", node.Data.T)
	fmt.Printf("区号：%s\n", node.Data.Q)
	fmt.Println()
}

func (node *Node)SetQ(q string){
	node.Data.Q = q
}

type f1 func(n *Node) interface{}

func Map(node *Node, f f1){
	if node == nil {
		return
	}
	Map(node.Left, f)
	f(node)
	Map(node.Right, f)
}

func Reduce(node *Node, f f1) int{
	var sum int
	if node == nil{
		return sum
	}
	sum += f(node).(int)

	left := node.Left
	right := node.Right
	return Reduce(left, f)+Reduce(right, f)+sum
}