func subarraySum(nums []int, k int) int {
	answer := 0
	for i:=0; i<len(nums); i++{
		currSum := 0
		for j:=i; j<len(nums); j++{
			currSum += nums[j]
			if currSum == k{
				answer++
			}
		}
	}
	return answer
}

