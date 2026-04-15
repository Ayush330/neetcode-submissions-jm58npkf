func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // Day 0 Initial States
    // If we start "holding" on day 0, we must have bought it.
    holding := -prices[0] 
    ready := 0
    cooldown := 0

    for i := 1; i < len(prices); i++ {
        // We use temporary variables so we don't use "today's" 
        // values to calculate "today's" other states.
        prevHolding := holding
        prevReady := ready
        prevCooldown := cooldown

        // Follow the arrows!
        holding = max(prevHolding, prevReady - prices[i])
        cooldown = prevHolding + prices[i]
        ready = max(prevReady, prevCooldown)
    }

    // At the end, the answer is the max money we have 
    // without a stock in our hand (Ready or Cooldown).
    return max(ready, cooldown)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}