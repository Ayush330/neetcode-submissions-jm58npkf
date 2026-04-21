func maxSlidingWindow(nums []int, k int) []int {
    deque := []int{}
    result := []int{}
    
    for i := 0; i < len(nums); i++ {
        // 1. RETIREMENT: Check if the current king is out of the window
        // i - k is the index that just fell off the cliff
        if len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }

        // 2. PRUNING: Kick out weaker/older contenders from the back
        for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        
        // 3. ADDING: Add current index to the line
        deque = append(deque, i)

        // 4. RECORDING: If window is full, write down the front (the max)
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    return result
}