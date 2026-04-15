func minDistance(word1 string, word2 string) int {
    l1 := len(word1)+1
    l2 := len(word2)+1
    dp := make([][]int, l1)
    for i:=0; i<l1; i++{
        dp[i] = make([]int, l2)
    }

    for i:=0; i<l1; i++{
        dp[i][0] = i
    }
    for i:=0; i<l2; i++{
        dp[0][i] = i
    }
    for i:=1; i<l1; i++{
        for j:=1; j<l2; j++{
            if word1[i-1] == word2[j-1]{
                dp[i][j] = dp[i-1][j-1]
                continue
            }
            // replacement
            replacement := dp[i-1][j-1]
            // deletion
            deletion := dp[i-1][j]
            // insertion
            insertion := dp[i][j-1]

            dp[i][j] = min(replacement, min(deletion, insertion))+1
        }
    }
    return dp[l1-1][l2-1]
}


func min(a, b int)int{
    if a > b{
        return b
    }
    return a
}