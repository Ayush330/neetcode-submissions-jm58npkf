func largestRectangleArea(heights []int) int {
    size := len(heights)
    left := make([]int, size)
    right := make([]int, size)
    stack := []int{}
    for i:=0; i<size; i++{
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i]{
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0{
            left[i] = -1
        }else{
            left[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
    stack = []int{}
    for i:=size-1; i>=0; i--{
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i]{
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0{
            right[i] = size
        }else{
            right[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }

    ans := 0
    for i:=0; i<size; i++{
        ans = max(ans, (right[i]-left[i]-1)*heights[i])
    }
    return ans
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
