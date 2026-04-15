func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    
    // Create a 2D table
    // dp[i][j] = cost to convert word1[:i] to word2[:j]
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }

    // --- BASE CASES (Initialization) ---
    
    // If word2 is empty (j=0), we must delete all chars from word1
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    
    // If word1 is empty (i=0), we must insert all chars from word2
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }

    // --- THE LOOPS ---
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            
            // Note: dp index 'i' corresponds to string index 'i-1'
            // because dp is 1-based (0 is empty string).
            if word1[i-1] == word2[j-1] {
                // MATCH: No cost. Just inherit from diagonal.
                dp[i][j] = dp[i-1][j-1]
            } else {
                // MISMATCH: 1 + min(Insert, Delete, Replace)
                
                insert  := dp[i][j-1] // Left
                delete  := dp[i-1][j] // Up
                replace := dp[i-1][j-1] // Diagonal
                
                dp[i][j] = 1 + min(insert, min(delete, replace))
            }
        }
    }

    return dp[m][n]
}

func min(a, b int) int {
    if a < b { return a }
    return b
}