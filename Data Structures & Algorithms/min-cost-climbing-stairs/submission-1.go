func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2{
		return min(cost[0], cost[1])
	}
    dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i:=2; i<len(cost); i++{
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	//return min(dp[len(dp)-1], dp[len(dp)-2])
	return minArr(dp[len(dp)-2:])
}

func minArr(arr []int)int{
	if arr[0] < arr[1]{
		return arr[0]
	}
	return arr[1]
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}