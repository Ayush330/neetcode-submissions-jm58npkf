func threeSum(nums []int) [][]int {
    answer := make([][]int, 0)
    if len(nums) < 3 {
        return answer
    }
    sort.Ints(nums) // Sorting is correct

    for i := 0; i <= len(nums)-3; i++ {
        // FIX 1: Skip Outer Duplicates (The Anchor)
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        target := -1 * nums[i]
        start := i + 1
        end := len(nums) - 1

        for start < end {
            currSum := nums[start] + nums[end]
            if currSum == target {
                // Note: Better to put nums[i] (which is -target) first for sorted order
                answer = append(answer, []int{nums[i], nums[start], nums[end]})
                
                // FIX 2: Move pointers AND skip inner duplicates
                start++
                end--
                
                // While start is valid and is same as previous, skip it
                for start < end && nums[start] == nums[start-1] {
                    start++
                }
                // While end is valid and is same as previous, skip it
                for start < end && nums[end] == nums[end+1] {
                    end--
                }

            } else if currSum > target {
                end--
            } else {
                start++
            }
        }
    }
    return answer
}