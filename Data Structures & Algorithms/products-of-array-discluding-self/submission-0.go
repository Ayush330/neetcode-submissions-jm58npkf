func productExceptSelf(nums []int) []int {
    size := len(nums)
    leftProduct := make([]int, size)
    leftProduct[0] = 1
    for i:=1; i<size; i++{
        leftProduct[i] = nums[i-1]*leftProduct[i-1]
    }

    rightProduct := make([]int, size)
    rightProduct[size-1] = 1
    for i:=size-2; i>=0; i--{
        rightProduct[i] = nums[i+1] * rightProduct[i+1]
    }

    result := make([]int, size)
    for i:=0;i<size;i++{
        result[i] = leftProduct[i]*rightProduct[i]
    }
    return result
}
