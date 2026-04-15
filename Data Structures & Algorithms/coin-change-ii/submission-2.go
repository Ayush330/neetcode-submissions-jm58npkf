func change(amount int, coins []int) int {
    dp := make([][]int, amount+1)
    for i:=0; i<amount+1; i++{
        dp[i] = make([]int, len(coins))
    }

    for j := 0; j < len(coins); j++ {
        dp[0][j] = 1
    }

    for i:=1; i<=amount; i++{
        for j:=0; j<len(coins); j++{
            x := 0
            if j > 0 {
                x = dp[i][j-1]
            }

            // Option 2: Include the current coin
            // We stay in the same column 'j' (unlimited supply), but reduce amount 'i'
            y := 0
            if i >= coins[j] {
                y = dp[i-coins[j]][j]
            }

            // Total ways = Exclude + Include
            dp[i][j] = x + y
        }
    }
    return dp[amount][len(coins)-1]
}
