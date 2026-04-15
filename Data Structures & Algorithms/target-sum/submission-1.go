func findTargetSumWays(nums []int, target int) int {
    return solve(0, nums, target, 0)
}




func solve(currIndex int, nums []int, target, running int)int{
    if currIndex >= len(nums){
        if target == running{
            return 1
        }
        return 0
    }

    return solve(currIndex+1, nums, target, running+nums[currIndex]) + solve(currIndex+1, nums, target, running-nums[currIndex])
}
