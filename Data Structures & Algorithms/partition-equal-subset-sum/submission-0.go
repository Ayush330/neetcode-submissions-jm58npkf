func canPartition(nums []int) bool {
    // two slices
    // a number can go in any slices
   sum := 0
   for _, num := range nums{
        sum+=num
   }
   if sum % 2 !=0{
        return false
   }
   half := sum/2
    return helper(0, nums, half, 0)
}

func helper(current int, nums []int, target, currSum int) bool{
    if currSum == target{
        return true
    }
    if current == len(nums){
        return false
    }
    return helper(current+1, nums, target, currSum+nums[current]) ||
            helper(current+1, nums, target, currSum)
}
