func maxProduct(nums []int) int {
    dp := make([][]int, len(nums)+1)
    for i := 0; i<=len(nums); i++{
        dp[i] = make([]int, len(nums)+1)
    }
    maxm := nums[0]
    for i:=1; i<=len(nums); i++{
        for j:=i; j<=len(nums); j++{
            if i == j{
                dp[i][j] = nums[i-1]
            }else{
                dp[i][j] = dp[i][j-1]*nums[j-1]
            }
            maxm = max(dp[i][j], maxm)
        }
    }
    return maxm
}

func max(a, b int)int{
    if a < b{
        return b
    }
    return a
}