func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	answer := 0
	curr := 0
	for i:=0; i<len(nums); i++{
		curr += nums[i]
		target := curr - k
		if freq, exists := mp[target]; exists{
			answer += freq
		}
		mp[curr]++
	}
	return answer
}
