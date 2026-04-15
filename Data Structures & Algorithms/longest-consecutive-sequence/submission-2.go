func longestConsecutive(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    sort.Ints(nums)
    fmt.Printf("%v", nums)
    maxm := 1
    curr := 1
    for i:=1;i<len(nums);i++{
        if nums[i]-nums[i-1] == 1{
            curr++
        }else if nums[i] == nums[i-1]{
            continue
        }else{
            curr = 1
        }
        maxm = max(curr, maxm)
    }
    return maxm
}

func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}