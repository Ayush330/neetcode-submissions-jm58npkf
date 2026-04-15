func canCompleteCircuit(gas []int, cost []int) int {
    // 1. Sanity Check (Global)
    totalGas, totalCost := 0, 0
    for i := 0; i < len(gas); i++ {
        totalGas += gas[i]
        totalCost += cost[i]
    }
    if totalGas < totalCost {
        return -1
    }

    // 2. Greedy Search (Local)
    currentTank := 0
    start := 0
    
    for i := 0; i < len(gas); i++ {
        // Add gas, subtract cost for THIS leg of the trip
        currentTank += gas[i] - cost[i]
        
        // If tank goes negative, we failed.
        if currentTank < 0 {
            // We cannot start at 'start', nor any station up to 'i'.
            // So start at i + 1.
            start = i + 1
            
            // Reset tank to 0 for the new attempt
            currentTank = 0
        }
    }
    
    return start
}