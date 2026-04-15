func maxArea(heights []int) int {
    // start from start and end, find the water that it can hold
    start := 0; end := len(heights)-1
    ans := 0
    for start < end{
        ans = max(ans, (end-start)*min(heights[start], heights[end]))
        //fmt.Println("Answer is: ", ans)
        shiftStart := heights[start]
        shiftEnd := heights[end]
        if shiftStart > shiftEnd {
            end--
        }else{
            start++
        }
    }
    return ans
}


func max(a, b int) int{
    if a < b{
        return b
    }
    return a 
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}