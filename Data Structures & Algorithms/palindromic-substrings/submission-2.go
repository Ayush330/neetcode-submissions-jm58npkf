func countSubstrings(s string) int {
    l := len(s)
    dp := make([][]bool, l)
    for i := range dp {
        dp[i] = make([]bool, l)
    }
    count := 0
    for i:=0; i<l; i++{
        dp[i][i] = true
        count++
    }

    for i:=0; i<l-1; i++{
        if s[i] == s[i+1]{
            dp[i][i+1] = true
            count++
        }
    }

    for i:=3; i<=l; i++{
        for j:=0; j<=l-i; j++{
            k := j+i-1

            if s[j] == s[k] && dp[j+1][k-1]{
                dp[j][k] = true
                count++
            }
        }
    }
    return count
}

