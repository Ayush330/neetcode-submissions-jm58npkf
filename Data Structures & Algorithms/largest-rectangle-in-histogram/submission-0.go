func largestRectangleArea(heights []int) int {
    // solve it using DP
    minh := make([][]int, len(heights))
    for i:=0; i<len(heights); i++{
        minh[i] = make([]int, len(heights))
    }


    for i:=0; i<len(heights); i++{
        for j:=i; j<len(heights); j++{
            if i==j{
                minh[i][j] = heights[i]
            }else{
                minh[i][j] = min(minh[i][j-1], heights[j])
            }
        }
    }

    ans := heights[0]
    for i:=0; i<len(heights); i++{
        for j:=i; j<len(heights); j++{
            ans = max((j-i+1)*minh[i][j], ans)
        }
    }
    return ans
}


func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}

func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}