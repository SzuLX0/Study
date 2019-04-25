package Map

type function func(a interface{}) interface{}

func Map(a interface{}, f function) interface{}{
	switch a.(type) {
	case int:
		return a
	case []int:
		b := make([]int, 0)
		for _, v:= range a.([]int){
			b = append(b, f(v).(int))
		}
		return b

	default:
		//return like: [[1], [1,2], [1,2,3]]
		m, ok := a.([]interface{})
		if ok == true{
			b := make([][]int, len(m))
			for i, v:= range m{
				b[i] = append(b[i], f(v).([]int)...)
			}
			return b
		}

	}
	return nil
}



