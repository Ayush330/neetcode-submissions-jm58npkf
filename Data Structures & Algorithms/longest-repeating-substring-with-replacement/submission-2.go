func characterReplacement(s string, k int) int {
	maxLength := 0
	for i:=0; i<len(s); i++{
		mp := make(map[byte]int)
		maxFreq := 0
		for j:=i; j<len(s); j++{
			mp[s[j]]++
			if mp[s[j]] > maxFreq{
				maxFreq = mp[s[j]]
			}
			replacementsReq := (j-i)+1 - maxFreq
			if replacementsReq <= k{
				maxLength = max(maxLength, (j-i)+1)
			}
		}
	}
	return maxLength
}


func max(a, b int)int{
	if a>b{
		return a
	}
	return b
}
