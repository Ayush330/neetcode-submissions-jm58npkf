func rob(nums []int) int {
    // start from 0 to l-1
	// start from 1 to l

	// take the maximum from the two I guess

	l := len(nums)

	if l == 1{
		return nums[l-1]
	}

	if l == 2{
		return max(nums[l-1], nums[l-2])
	}


	//dp := make([]int, l)

	var answer int
	//dp[0] = nums[0]
	//dp[1] = max(nums[0], nums[1])

	first, second := nums[0], max(nums[0], nums[1])

	for i:=2; i<l-1; i++{
		//dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		first, second = second, max(second, first+nums[i])
	}

	answer = second //dp[l-2]

	// dp[0] = 0
	// dp[1] = nums[1]
	// dp[2] = max(nums[2], nums[1])

	first, second = nums[1], max(nums[1], nums[2])
	for i:=3; i<l; i++{
		//dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		first, second = second, max(second, first+nums[i])
	}


	return max(second, answer)
}
