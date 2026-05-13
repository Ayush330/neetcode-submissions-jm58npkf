func rob(nums []int) int {
    // start from 0 to l-1
	// start from 1 to l

	// take the maximum from the two I guess

	l := len(nums)

	if l == 1{
		return nums[l-1]
	}

	if l == 2{
		return max(nums[l-1], nums[l-2])
	}

	return max(helper(nums[:l-1]), helper(nums[1:]))
}


func helper(nums []int)int{
	first, second := nums[0], max(nums[0], nums[1])
	for i:=2; i<len(nums); i++{
		first, second = second, max(second, first+nums[i])
	}
	return second
}
