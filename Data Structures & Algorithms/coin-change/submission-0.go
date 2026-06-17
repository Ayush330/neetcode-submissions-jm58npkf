func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
	globalMax := amount + 1
	for i:=0; i<len(dp); i++{
		dp[i] = globalMax
	}
	dp[0] = 0
	for i:=0; i<len(dp); i++{
		for j:=0; j<len(coins); j++{
			coin := coins[j]
			if i-coin >= 0{
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == globalMax{
		return -1
	}
	return dp[amount]
}

func min(a, b int)int{
	if a<b{
		return a
	}
	return b
}
