func maxProfit(prices []int) int {
   l := len(prices)
    if l < 2{
        return 0
    }

    dp := make([][]int, l)
    for i:=0; i<l; i++{
        dp[i] = make([]int, 3)
    }

    dp[0][0] = -prices[0]
    for i:=1; i<l; i++{
        // handle the case when he held the stock
        dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
        // he sold 
        dp[i][1] = dp[i-1][0] + prices[i]
        // cooldown
        dp[i][2] = max(dp[i-1][1], dp[i-1][2])
    }


    return max(dp[l-1][0], max(dp[l-1][1], dp[l-1][2]))
}


func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}
