func trap(height []int) int {
	right := make([]int, len(height))
	right[len(height)-1] = height[len(height)-1]
	for i:=len(height)-2; i>=0; i--{
		right[i] = max(height[i], right[i+1])
	}

	ans := 0
	leftMax := height[0]
	for i:=0; i<len(height); i++{
		ans += max(0, min(leftMax, right[i]) - height[i])
		leftMax = max(leftMax, height[i])
	}
	return ans
}

func min(a, b int)int{
	if a<b{
		return a 
	}
	return b
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
