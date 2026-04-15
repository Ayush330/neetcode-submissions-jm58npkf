func dailyTemperatures(t []int) []int {
	stack := []int{}
	ans := make([]int, len(t))
	i := 0
	for i<len(t){
		if len(stack) == 0{
			stack = append(stack, i)
			i++
			continue
		}
		for len(stack) > 0 && t[stack[len(stack)-1]] < t[i]{
			ans[stack[len(stack)-1]] = i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
		i++
	}
	return ans
}
