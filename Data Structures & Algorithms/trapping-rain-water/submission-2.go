func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	// find the left maximum
	left[0] = height[0]
	for i:=1; i<len(height); i++{
		left[i] = max(left[i-1], height[i])
	}

	// find the right maximum
	right[len(height)-1] = height[len(height)-1]
	for i:=len(height)-2; i>=0; i--{
		right[i] = max(right[i+1], height[i])
	}
	answer := 0
	for i:=1; i<len(height)-1; i++{
		answer += max(min(left[i], right[i]) - height[i], 0)
	}
	return answer
}

func min(a, b int)int{
	if a<b{
		return a
	}
	return b
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
