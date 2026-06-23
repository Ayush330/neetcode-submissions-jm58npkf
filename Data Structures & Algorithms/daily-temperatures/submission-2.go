func dailyTemperatures(t []int) []int {
	stack := make([]int, 0)

	answer := make([]int, len(t))
	stack = append(stack, 0)

	for i:=1; i<len(t); i++{
		for len(stack) > 0 && t[stack[len(stack)-1]] < t[i]{
			answer[stack[len(stack)-1]] = i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return answer

}
