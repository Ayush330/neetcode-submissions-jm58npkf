func jump(nums []int) int {
	if len(nums) == 1{
		return 0
	}
	steps := 0
	currEnd := 0
	nextMax := 0
	for i:=0; i<len(nums)-1; i++{
		if i == currEnd{
			nextMax = max(nextMax, i+nums[i])
			steps++
			currEnd = nextMax
			nextMax = 0
		}else{
			nextMax = max(nextMax, i+nums[i])
		}

	}
	return steps
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
