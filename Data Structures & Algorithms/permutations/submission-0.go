func permute(nums []int) [][]int {
	answer := [][]int{}
	visited := make([]bool, len(nums))
	solve(nums, []int{}, &answer, visited)
	return answer
}


func solve(nums []int, curr []int, ans *[][]int, visited []bool){
	if len(curr) == len(nums){
		tmp := make([]int, len(nums))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
		return 
	}

	for j:=0; j<len(nums); j++{
		if visited[j]{
			continue
		}
		visited[j] = true
		curr = append(curr, nums[j])
		solve(nums, curr, ans, visited)
		curr = curr[:len(curr)-1]
		visited[j] = false
	}
	return
}