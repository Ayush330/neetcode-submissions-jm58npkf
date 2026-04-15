func minEatingSpeed(piles []int, h int) int {
    start := 1
    end := start
    for i:=0; i<len(piles); i++{
        end = max(end, piles[i])
    }
    answer := end
    for start <= end{
        mid := start + (end-start)/2
        tt := timeTaken(piles, mid)
        if tt <= h{
            answer = min(mid, answer)
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return answer
}


func min(a, b int)int{
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


func timeTaken(piles []int, k int)int{
    time := 0
    for _, count := range piles{
        abs := count/k
        rem := 0
        if count%k == 0{
            rem = 0
        }else{
            rem = 1
        }
        time = time +  abs + rem
    }
    return time
}