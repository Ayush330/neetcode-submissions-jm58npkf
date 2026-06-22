import s "slices"

func combinationSum2(candidates []int, target int) [][]int {
	s.SortFunc(candidates, func(a, b int)int{
		return a-b
	})
	answer := make([][]int, 0)
	solve(candidates, 0, 0, target, []int{}, &answer)
	return answer
}

func solve(nums []int, i, sum, target int, curr []int, ans *[][]int){
	if sum == target{
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return 
	}

	if sum > target || i >= len(nums){
		return
	}

	for j:=i; j<len(nums); j++{
		if j>i && nums[j] == nums[j-1]{
			continue
		}
		curr = append(curr, nums[j])
		solve(nums, j+1, sum+nums[j], target, curr, ans)
		curr = curr[:len(curr)-1]
	}
	return 
}