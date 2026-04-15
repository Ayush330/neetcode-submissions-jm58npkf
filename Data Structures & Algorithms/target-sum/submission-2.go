func findTargetSumWays(nums []int, target int) int {
    totalSum := 0
    for _, n := range nums {
        totalSum += n
    }

    // 1. Math Check: Can we split the array into two integers?
    // derived from: x = (target + totalSum) / 2
    if (totalSum+target)%2 != 0 || totalSum < abs(target) {
        return 0
    }

    bagSize := (totalSum + target) / 2
    if bagSize < 0 { return 0 } // Edge case if target is very negative

    // dp[j] = number of ways to fill a bag of capacity j
    dp := make([]int, bagSize+1)
    
    // 2. Base Case
    dp[0] = 1 

    // 3. The Loop
    for _, num := range nums {
        // BRUTAL CHECK: Iterate BACKWARDS.
        // If you go forwards (0 to bagSize), you solve Coin Change 2 (infinite supply).
        // You must go backwards to ensure 'num' is used only once per subset.
        for j := bagSize; j >= num; j-- {
            dp[j] = dp[j] + dp[j-num]
        }
    }

    return dp[bagSize]
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}