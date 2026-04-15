func findDuplicate(nums []int) int {
    // Phase 1: Finding the intersection point in the cycle
    tortoise := nums[0]
    hare := nums[0]

    for {
        tortoise = nums[tortoise]          // Move 1 step
        hare = nums[nums[hare]]             // Move 2 steps
        if tortoise == hare {
            break
        }
    }

    // Phase 2: Finding the entrance to the cycle (the duplicate)
    // Reset one pointer to the start
    ptr1 := nums[0]
    ptr2 := tortoise // Keep the other at the intersection point

    for ptr1 != ptr2 {
        ptr1 = nums[ptr1] // Both move at 1 step now
        ptr2 = nums[ptr2]
    }

    return ptr1 // This is the duplicate!
}