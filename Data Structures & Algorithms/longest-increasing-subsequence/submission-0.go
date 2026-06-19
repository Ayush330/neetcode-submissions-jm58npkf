func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
	for i:=0; i<len(nums); i++{
		dp[i] = 1
	}
	globalMax := 1
	for i:=0; i<len(nums); i++{
		// picking the num
		for j:=i; j>=0; j--{
			// can it be a part of the 
			if nums[i] > nums[j]{
				dp[i] = max(dp[i], dp[j]+1)
			}else{
				dp[i] = dp[i]
			}
			globalMax = max(globalMax, dp[i])
		}
	}
	return globalMax
}
