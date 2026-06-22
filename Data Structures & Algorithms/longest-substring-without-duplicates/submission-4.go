func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)

    start := 0
    maxLen := 0

    for end := 0; end < len(s); end++ {
        if idx, ok := mp[s[end]]; ok && idx >= start {
            start = idx + 1
        }

        mp[s[end]] = end
        maxLen = max(maxLen, end-start+1)
    }

    return maxLen
}