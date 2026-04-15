func canJump(nums []int) bool {
	cmax := 0
	for i:=0; i<len(nums); i++{
		if cmax >= len(nums)-1{
			return true
		}
		if cmax < i{
			return false
		}
		cmax = max(cmax, i+nums[i])
	}
	return false
}



func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}









