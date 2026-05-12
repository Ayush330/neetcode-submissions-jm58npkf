func climbStairs(n int) int {
    // dp[x] = dp[x-1] + dp[x-2]

	if n <= 2{
		return n
	}
	first := 1
	second := 2

	for i:=3; i<=n; i++{
		first, second = second, second + first
	}

	return second
}
