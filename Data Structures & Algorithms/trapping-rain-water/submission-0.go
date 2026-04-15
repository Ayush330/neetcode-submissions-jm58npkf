func trap(height []int) int {
    size := len(height)
    left := make([]int, size)
    right := make([]int, size)

    // populate the left
    left[0] = height[0]
    for i:=1; i<size;i++{
        left[i] = max(left[i-1], height[i])
    }

    right[size-1] = height[size-1]
    for i:=size-2; i>=0; i--{
        right[i] = max(right[i+1], height[i])
    }

    ans := 0
    for i:=1; i<size-1;i++{
        ans += max(min(right[i], left[i]) - height[i], 0)
    }
    return ans
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}
