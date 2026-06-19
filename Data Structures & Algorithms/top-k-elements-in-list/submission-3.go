import "slices"

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)

	for i:=0; i<len(nums); i++{
		mp[nums[i]]++
	}

	type m struct{
		key int
		val int
	}
	arr := make([]m, 0)
	for key, val := range mp{
		arr = append(arr, m{key, val})
	}

	slices.SortFunc(arr, func(a, b m)int {
		return b.val - a.val
	})

	answer := make([]int, k)
	for i:=0; i<k; i++{
		answer[i] = arr[i].key
	}
	return answer
}
