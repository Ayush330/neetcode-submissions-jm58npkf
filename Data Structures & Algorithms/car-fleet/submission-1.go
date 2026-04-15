import "slices"

type pair struct{
	posn int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	l := len(speed)
	arr := make([]pair, l, l)
	for i:=0; i<l ; i++{
		arr[i] = pair{position[i], speed[i]}
	}
	slices.SortFunc(arr, func(a, b pair)int{
		if a.posn == b.posn{
			return -1*(a.speed - b.speed)
		}
		return a.posn-b.posn
	})

	stack := []float64{}
	for i:=l-1; i>-1; i--{
		t := float64(target - arr[i].posn) / float64(arr[i].speed)
		if len(stack) == 0{
			stack = append(stack, t)
			continue
		}
		top := stack[len(stack)-1]
		if t > top{
			stack = append(stack, t)
		}
	}
	return len(stack)
}
