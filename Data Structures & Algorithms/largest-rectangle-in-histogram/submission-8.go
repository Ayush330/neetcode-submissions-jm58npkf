func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := []int{} // Start with a completely empty waiting room
	ans := 0

	for index, h := range heights {
		
		// 2. The Condition: Must be AND (&&) to prevent panics!
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			
			// 3. Pop the tall guy out FIRST so we don't infinite loop!
			tallBarIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] 
			
			tallBarHeight := heights[tallBarIndex]
			
			// 4. The Ruler: Calculate the width
			width := 0
			if len(stack) == 0 {
				// If the room is empty, it means there is no left wall.
				// The rectangle stretches all the way from the start to the current index.
				width = index
			} else {
				// The width is sandwiched between the current guy (index) 
				// and the guy newly sitting at the top of the stack.
				leftWallIndex := stack[len(stack)-1]
				width = index - leftWallIndex - 1
			}
			
			// Calculate area and update the global max
			ans = max(ans, width * tallBarHeight)
		}
		
		// 5. The new guy safely enters the waiting room
		stack = append(stack, index)
	}
	return ans
}
