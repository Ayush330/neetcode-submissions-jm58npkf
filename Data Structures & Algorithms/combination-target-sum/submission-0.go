func combinationSum(nums []int, target int) [][]int {
    answer := make([][]int, 0)
	solve(nums, 0, 0, target, []int{}, &answer)
	return answer
}


func solve(nums []int, i, sum, target int, curr []int, global *[][]int){
	if sum == target{
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*global = append(*global, tmp)
		return
	}
	if i >= len(nums) || sum > target{
		return
	}

	for j:=i; j<len(nums); j++{
		curr = append(curr, nums[j])
		solve(nums, j, sum + nums[j], target, curr, global)
		curr = curr[:len(curr)-1]
	}
	return
}