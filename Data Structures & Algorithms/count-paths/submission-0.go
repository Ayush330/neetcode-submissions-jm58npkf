func uniquePaths(m int, n int) int {
    // move[i][j] = move[i-1][j] + move[i][j-1]
    // move[0][0] = 0
    // move[i][0], move[0][i] = 1


    dp := make([][]int, m)
    for i:=0; i<m; i++{
        dp[i] = make([]int, n)
    }

    for i:=0; i<m; i++{
        dp[i][0] = 1
    }

    for j:=0; j<n; j++{
        dp[0][j] = 1
    }

    for i:=1; i<m; i++{
        for j:=1; j<n; j++{
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}
