func canJump(nums []int) bool {
    maxJump := 0
	for i:=0; i<len(nums); i++{
		if maxJump >= i{
			maxJump = max(maxJump, i+nums[i])
		}else{
			return false
		}
	}
	return true
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
