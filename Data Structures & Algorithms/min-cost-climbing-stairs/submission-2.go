func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2{
		return min(cost[0], cost[1])
	}
    //dp := make([]int, len(cost))
	//dp[0] = cost[0]
	//dp[1] = cost[1]
	first, second := cost[0], cost[1]
	for i:=2; i<len(cost); i++{
		//dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		first, second = second, min(first, second)+cost[i]
	}
	//return min(dp[len(dp)-1], dp[len(dp)-2])
	return min(first, second)
}


func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}