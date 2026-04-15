func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    stack := []int{} // This will store indices

    for i := 0; i < n; i++ {
        // While the stack is not empty AND 
        // the current temp is warmer than the temp at the top of the stack
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            // Pop the index from the top
            prevIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            // Calculate the distance
            ans[prevIndex] = i - prevIndex
        }
        // Push the current index onto the stack
        stack = append(stack, i)
    }
    
    return ans
}