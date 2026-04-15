func largestRectangleArea(heights []int) int {
    stack := []int{-1}
	answer := 0
	for i:=0; i<=len(heights); i++{
		currH := 0
		if i < len(heights){
			currH = heights[i]
		}
		for len(stack) > 1 && currH < heights[stack[len(stack)-1]]{
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			right := i
			width := right - left - 1
			answer = max(answer, height * width)
		}
		stack = append(stack, i)
	}
	return answer
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
