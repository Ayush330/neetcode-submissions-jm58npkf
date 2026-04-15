func minEatingSpeed(piles []int, h int) int {
    minm := 1
    // for _, pile := range piles{
    //     maxm+=pile
    // }
    maxm := piles[0]
    for _, pile := range piles{
        maxm = max(maxm, pile)
    }
    mid := 0
    ans := maxm
    for maxm >= minm{
        mid = minm + (maxm-minm)/2
        timetaken := timeTaken(piles, mid)
        if timetaken > h{
            minm = mid+1
        }else {
            ans = min(mid, ans)
            maxm = mid-1
        }
    }
    return ans
}


func min(a, b int)int{
    if a<b{
        return a
    }
    return b
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func timeTaken(piles []int, speed int)int{
    timeTaken := 0
    for _, pile := range piles{
        offset := 0
        if pile%speed != 0{
            offset = 1
        }
        time := pile/speed
        timeTaken += (offset+time)
    }
    return timeTaken
}