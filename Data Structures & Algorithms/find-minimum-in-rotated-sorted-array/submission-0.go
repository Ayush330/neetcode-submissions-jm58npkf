func findMin(nums []int) int {
    size := len(nums)
    left := 0; right := size-1
    for left < right{
        mid := left + (right-left)/2
        if nums[mid] > nums[right]{
            left = mid+1
        }else{
            right = mid
        }
    }
    fmt.Println(left, right)
    return nums[right]
}
