func longestPalindrome(s string) string {
    n := len(s)
	if n < 2{
		return s
	}

	dp := make([][]bool, n)
	for i:=0; i<n; i++{
		dp[i] = make([]bool, n)
	}

	for i:=0; i<n; i++{
		dp[i][i] = true
	}

	start := 0
	maxLen := 1
	for i:=0; i<n-1; i++{
		if s[i] == s[i+1]{
			start = i
			maxLen = 2
			dp[i][i+1] = true
		}
	}

	for l:=3; l<=n; l++{
		for i:=0; i<=n-l; i++{
			j := l+i-1
			if s[i] == s[j]{
				if dp[i+1][j-1]{
					dp[i][j] = true
					if l > maxLen{
						maxLen = l
						start  = i
					}
				} 
			}
		}
	}
	fmt.Println("Start Is: ", start, " MaxLen Is: ", maxLen)
	return s[start:start+maxLen]
}

