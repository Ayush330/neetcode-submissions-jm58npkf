func largestRectangleArea(heights []int) int {
    stack := []int{-1}
    i:=0
    answer := 0
    heights = append(heights, -1)
    for len(stack)>0 && i<len(heights){
        curr :=  heights[i]
        //fmt.Println("data Is: ", stack[len(stack)-1])
        for len(stack) > 1 && curr < heights[stack[len(stack)-1]]{
            
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            left := stack[len(stack)-1]
            right := i
            area := (right-left-1)*height
            answer = max(answer, area)
        }
        stack = append(stack, i)
        i++
    }
    return answer
}

func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}
