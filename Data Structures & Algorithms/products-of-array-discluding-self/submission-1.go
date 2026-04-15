func productExceptSelf(nums []int) []int {
    size := len(nums)
    result := make([]int, size)
    result[0] = 1
    for i:=1; i<size; i++{
        result[i] = nums[i-1]*result[i-1]
    }

    rightVal := 1
    for i:=size-1; i>=0; i--{
        result[i] = result[i] * rightVal
        rightVal = rightVal * nums[i]
    }
    return result
}
