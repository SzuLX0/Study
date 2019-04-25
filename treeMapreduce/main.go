package main

import (
	"fmt"
	"github.com/szulx0/treeMapreduce/treeMapreduce"
)

var q  = make(map[string]string)

func main() {

	q["深圳"] = "755"

	var n1 = &treeMapreduce.Node{
		Data:treeMapreduce.Data{0, "小明", "深圳", "26530001", ""},
		Left:  nil,
		Right: nil,
	}
	n1.Left = &treeMapreduce.Node{
		Data:treeMapreduce.Data{Flag:0, N:"小王", A:"广州", T:"26530002", Q:""},
		Left:nil,
		Right:nil,
	}
	n1.Left.Left = &treeMapreduce.Node{
		Data:treeMapreduce.Data{Flag:0, N:"小丽", A:"深圳", T:"26530003", Q:""},
		Left:nil,
		Right:nil,
	}
	n1.Left.Right = &treeMapreduce.Node{
		Data:treeMapreduce.Data{Flag:0, N:"小红", A:"北京", T:"26530004", Q:""},
	}


	n1.MiddleOrder()
	//使用Map设置每个节点的区号
	treeMapreduce.Map(n1, f)
	n1.MiddleOrder()

	//使用map将深圳地区的flag设为1
	treeMapreduce.Map(n1, f1)
	//使用reduce计算深圳号码个数
	res := treeMapreduce.Reduce(n1, func(n *treeMapreduce.Node)interface{} {
		if n!=nil{
			return n.Flag
		}else{
			return 0
		}

	})
	fmt.Printf("深圳号码个数：%d\n", res)
}


//自动设置深圳的区号
func f(n *treeMapreduce.Node) interface{}{
	if n.A=="深圳"{
		n.SetQ(q["深圳"])
		fmt.Printf("设置%s的深圳区号\n", n.N )
		fmt.Println()
	}
	return nil
}

func f1(node *treeMapreduce.Node)interface{}{
	if node.A=="深圳"{
		node.Flag = 1
	}
	return nil
}
