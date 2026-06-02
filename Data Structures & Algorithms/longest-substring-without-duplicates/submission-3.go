func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int, 26)

	start, end := 0, 0
	ans := 0
	for end < len(s){
		char := s[end]
		if lastIndex, ok := mp[char]; ok{
			if lastIndex >= start{
				start = lastIndex+1
			}
		}
		mp[char] = end
		end++
		ans = max(ans, end-start)
	}
	return ans
}

func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}
