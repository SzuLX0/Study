package main

import (
	"fmt"
)

var count int = 0

func main() {
	var n int
	var a [][]int
	//var q = make([]int, len(a))

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		b := make([]int, n)
		a = append(a, b)
	}

	Queen(a)

}

func show(a [][]int) {
	count++
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}



func Queen(a [][]int) {
	if settleQueen(a, 0) == false {
		fmt.Println("There is no result")
	}
	fmt.Printf("递归结果：The sum of result is %d \n", count)
	count = 0
	fmt.Println()
	settleQueen2(a)
	fmt.Printf("循环结果：The sum of result is %d \n", count)
}

func settleQueen(a [][]int, y int) bool {
	N := len(a)
	if y == N {
		show(a)
		return true
	}

	res := false
	for i := 0; i < N; i++ {
		if check(a, i, y) {
			a[i][y] = 1
			//make res true if has any result
			res = settleQueen(a, y+1) || res
			//backtrace
			a[i][y] = 0
		}
	}

	return res
}

func settleQueen2(a [][]int) {
	n := len(a)
	var i int
	var j int
	var flag int
	for i < n {
		flag = -1
		//找到第i行皇后的位置
		for j < n {
			if check(a, i, j) {
				a[i][j] = 1
				flag = j
				j = 0
				//fmt.Println(check(a, i, j))
				break
			} else {
				j++
			}
		}

		//若该行不能放置皇后，回溯
		if flag == -1 {
			if i == 0 {
				break
			} else {
				i--
				for k := 0; k < n; k++ {
					if a[i][k] == 1 {
						j = k + 1
						a[i][k] = 0
						break
					}
				}
				continue
			}
		}

		//当在最后一行找到，列移动寻找下一个解
		if i == n-1 {
			show(a)
			a[i][flag] = 0
			j = flag + 1
			continue
		}
		i++

	}

}


func check(a [][]int, x, y int) bool{
	length := len(a)

	for i:=0; i<length; i++{
		if a[x][i] == 1 || a[i][y]==1{
			return false
		}
		//检查左上角是否存在皇后
		if x-i>=0 && y-i>=0 && a[x-i][y-i]==1{
			return false
		}
		//检查右上角
		if x-i>=0 && y+i<length && a[x-i][y+i]==1{
			return false
		}
		//左下
		if x+i<length && y-i>=0 && a[x+i][y-i]==1{
			return false
		}
		//右下
		if x+i<length && y+i<length && a[x+i][y+i]==1{
			return false
		}
	}
	return true
}

