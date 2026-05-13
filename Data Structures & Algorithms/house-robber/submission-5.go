func rob(nums []int) int {
    // (dp[i] = dp[i-1], dp[i-2]+nums[i])
	l := len(nums)

	if l == 1{
		return nums[l-1]
	}

	if l == 2{
		return max(nums[l-1], nums[l-2])
	}

	dp := make([]int, l)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	
	for i:=2; i<l; i++{
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[l-1]
}


func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
