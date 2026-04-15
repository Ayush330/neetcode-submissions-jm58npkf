func trap(height []int) int {
    size := len(height)
    if size < 3 { return 0 }

    leftMax := height[0]
    rightMax := height[size-1]
    ans := 0
    i := 1; j := size-2

    for i <= j {
        if leftMax < rightMax {
            // Process the bar at index 'i'
            // Water = current bottleneck - current height
            if leftMax > height[i] {
                ans += leftMax - height[i]
            }
            // Update the wall for the NEXT step
            leftMax = max(leftMax, height[i])
            i++ // Now move to the next bar
        } else {
            // Process the bar at index 'j'
            if rightMax > height[j] {
                ans += rightMax - height[j]
            }
            rightMax = max(rightMax, height[j])
            j--
        }
    }
    return ans
}