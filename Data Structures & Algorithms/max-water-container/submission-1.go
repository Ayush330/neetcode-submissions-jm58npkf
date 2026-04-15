func maxArea(heights []int) int {
    start := 0; end := len(heights)-1
    answer := 0
    for start < end{
        answer = max(answer, (end-start)*min(heights[start], heights[end]))
        if heights[start] < heights[end]{
            start++
        }else{
            end--
        }
    }
    return answer
}


func max(a, b int)int{
    if a>b{
        return a 
    }
    return b
}

func min(a, b int)int{
    if a<b{
        return a 
    }
    return b
}
