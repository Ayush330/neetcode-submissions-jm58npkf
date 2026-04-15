func characterReplacement(s string, k int) int {
	maxFreq := 0
	maxLength := 0
	mp := make(map[byte]int)

	left:=0
	for right:=0; right<len(s); right++{
		mp[s[right]]++
		maxFreq = max(maxFreq, mp[s[right]])

		for right-left+1 - maxFreq > k{
			mp[s[left]]--
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
