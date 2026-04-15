func rob(nums []int) int {
    if len(nums)==1{
		return nums[0]
	}

	if len(nums) == 2{
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	// first iteration
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i:=2; i<len(nums)-1;i++{
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	curr := dp[len(nums)-2]
	// second iteration
	dp[0] = 0
	dp[1] = nums[1]
	dp[2] = max(nums[1], nums[2])
	for i:=3; i<len(nums); i++{
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return max(curr, dp[len(dp)-1])
}



func max(a, b int)int{
	if a<b{
		return b
	}
	return a
}