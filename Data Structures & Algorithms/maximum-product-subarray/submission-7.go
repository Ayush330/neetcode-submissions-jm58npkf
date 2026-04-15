func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // Initialize with the first element
    maxSoFar := nums[0]
    minSoFar := nums[0]
    globalMax := nums[0]

    for i := 1; i < len(nums); i++ {
        curr := nums[i]

        // We must calculate the new candidates before updating maxSoFar,
        // because updating maxSoFar first would ruin the calculation for minSoFar.
        prod1 := curr * maxSoFar
        prod2 := curr * minSoFar

        // New maxSoFar is the best of (itself, current*max, current*min)
        maxSoFar = max(curr, max(prod1, prod2))
        
        // New minSoFar is the worst of (itself, current*max, current*min)
        minSoFar = min(curr, min(prod1, prod2))

        // Update the overall best result we've ever seen
        globalMax = max(globalMax, maxSoFar)
    }

    return globalMax
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}