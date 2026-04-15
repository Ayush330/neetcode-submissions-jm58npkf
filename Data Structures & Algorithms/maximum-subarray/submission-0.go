func maxSubArray(nums []int) int {
    globalMax := nums[0]
    currMax := nums[0]
    for i:=1; i<len(nums); i++{
        tmp := currMax + nums[i]
        if tmp >= nums[i]{
            currMax = tmp
        }else{
            currMax = nums[i]
        }
        globalMax = max(globalMax, currMax)
    }
    return globalMax
}


func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}