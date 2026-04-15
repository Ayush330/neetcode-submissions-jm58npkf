func longestConsecutive(nums []int) int {
    // sort the nums array
    if len(nums) == 0{
        return 0
    }
    sort.Slice(nums, func(i, j int)bool{
        return nums[i] <= nums[j]
    })
    maxm := 1
    answer := 1
    for i:=1;i<len(nums); i++{
        if nums[i] == nums[i-1]{
            continue
        }
        if nums[i]-nums[i-1] == 1{
            answer++
        }else{
            maxm = max(answer, maxm)
            answer = 1
        }
    }
    maxm = max(answer, maxm)
    return maxm
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
