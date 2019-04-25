package main

import (
	"fmt"
	"github.com/szulx0/mapReduce/Map"
	"github.com/szulx0/mapReduce/Reduce"
)


func main() {
	var a []int
	var n int
	var tmp int
	var b []int
	var res int

	//1. 计算平方和
	fmt.Println("计算平方和")
	fmt.Println("请输入数组的大小")
	fmt.Scanf("%d", &n)
	a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scanf("%d", &tmp)
		a = append(a, tmp)
	}
	//利用Map将数组转化成原来的平方

	b = Map.Map(a, func(a interface{}) interface{} {
		return a.(int)*a.(int)
	}).([]int)
	res = Reduce.Reduce(b, func(a, b interface{}) interface{}{return a.(int)+b.(int)}).(int)
	fmt.Printf("该数组的平方和为： %d\n", res)

//--------------------------------------------------------------------------

	//2. 计算数组中正数个数
	fmt.Println("计算数组中正数的个数")
	fmt.Println("请输入数组的大小")
	fmt.Scanf("%d", &n)
	a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scanf("%d", &tmp)
		a = append(a, tmp)
	}


	//使用map将数组中正数设为1， 非正数设为0
	b = Map.Map(a, func(i interface{}) interface{} {
		if i.(int)>0{
			return 1
		}else{
			return 0
		}
	}).([]int)

	//利用reduce计算转换后数组的和即为数组中正数个数
	res = Reduce.Reduce(b, func(a, b interface{}) interface{}{return a.(int)+b.(int)}).(int)
	fmt.Printf("该数组中正数个数为： %d\n", res)

//-------------------------------------------------------------------------------
	//3. 数组合并
	fmt.Println("数组的数组")
	a1 := []interface{}{0, 8, 9}
	b1 := []interface{}{1, 2, 3}

	c1 := make([]interface{}, 4)
	for i:=0; i<4; i++{
		c1[i] = i*i
	}
	a1 = append(a1, c1)
	a1 = append(a1, b1)
	c1 = append(c1, a1)
	a1 = append(a1, c1)
	fmt.Println(a1)

	//使用Map将数组中每一项展开成一维数组
	result := Map.Map(a1, flatten).([][]int)
	b = Reduce.Reduce(result, func(a,b interface{})interface{}{
		return append(a.([]int), b.([]int)...)}).([]int)
	fmt.Println(b)

}



func flatten(a interface{}) interface{}{
	res := make([]int, 1)
	switch a.(type) {
	case int:
		res[0] = a.(int)
		return res
	case []int:
		return a.([]int)
	default:
		return Reduce.Reduce(Map.Map(a, flatten), func(a,b interface{})interface{}{
			return append(a.([]int), b.([]int)...)})
	}
}

