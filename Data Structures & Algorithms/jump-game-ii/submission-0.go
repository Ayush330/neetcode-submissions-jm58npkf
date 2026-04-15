func jump(nums []int) int {
    l, r := 0, 0
    maxm := 0
    count := 0
    for r < len(nums)-1{
        for i:=l; i<=r; i++{
            maxm = max(maxm, i+nums[i])
        }
        r = max(r, maxm)
        count++
    }
    return count
}

func maxm(a, b int)int{
    if a > b{
        return b
    }
    return a
}
