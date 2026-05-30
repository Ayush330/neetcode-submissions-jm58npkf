func largestRectangleArea(h []int) int {
    stack := make([]int, 0, len(h))
    var ans int
    
    // Loop through all elements, PLUS ONE extra iteration at the end
    // to force the stack to completely empty out.
    for i := 0; i <= len(h); i++ {
        
        // If we are at the end, treat the current height as 0 to force pops
        currentHeight := 0
        if i < len(h) {
            currentHeight = h[i]
        }
        
        // POP CONDITION: While stack has items AND the current bar is a "wall" 
        // (shorter than the bar at the top of the stack)
        for len(stack) > 0 && currentHeight < h[stack[len(stack)-1]] {
            // 1. Pop the top element
            poppedIdx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            // 2. Calculate the area for the popped element
            height := h[poppedIdx]
            width := 0
            
            if len(stack) == 0 {
                // If stack is empty, it means this bar was the shortest one 
                // we've seen so far, so it stretches all the way back to index 0.
                width = i 
            } else {
                // The width is bounded by the current index (right wall) 
                // and the NEW top of the stack (left wall)
                newTopIdx := stack[len(stack)-1]
                width = i - newTopIdx - 1
            }
            
            // 3. Update Max Area
            ans = max(ans, height * width)
        }
        
        // Unconditionally push the current index
        stack = append(stack, i)
    }
    
    return ans
}

func max(a, b int) int {
    if a > b { return a }
    return b
}