package Reduce

type function func(a, b interface{})interface{}

func Reduce(a interface{}, f function) interface{}{
	switch a.(type) {
	case int:
		return a.(int)
	case []int:
		a1 := a.([]int)
		if len(a1)<=1{
			return a1[0]
		}else if len(a1) == 2{
			return f(a1[0], a1[1])
		}
		b := a1[0]
		for i:=1; i<len(a1); i++{
			b = f(b, a1[i]).(int)
		}
		return b

	case [][]int:
		a2 := a.([][]int)
		if len(a2)<=1{
			return a2[0]
		}else if len(a2) == 2{
			return f(a2[0], a2[1])
		}
		b := a2[0]
		for i:=1; i<len(a2); i++{
			b = f(b, a2[i]).([]int)
		}
		return b
	}
	return nil
}
