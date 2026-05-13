func longestPalindrome(s string) string {
    l := len(s)
    if l <= 1 {
        return s
    }

    startIndex, maxLength := 0, 1
    dp := make([][]bool, l)
    for i := range dp {
        dp[i] = make([]bool, l)
        dp[i][i] = true // Base case: length 1
    }

    // Step 1: Check length 2
    for i := 0; i < l-1; i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = true
            startIndex = i
            maxLength = 2
        }
    }

    // Step 2: Check length 3 and above
    // We iterate by 'length' to ensure subproblems are solved first
    for length := 3; length <= l; length++ {
        for i := 0; i <= l-length; i++ {
            j := i + length - 1 // End index
            
            // The DP Magic: Ends match AND middle is a palindrome
            if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
                startIndex = i
                maxLength = length
            }
        }
    }

    return s[startIndex : startIndex+maxLength]
}