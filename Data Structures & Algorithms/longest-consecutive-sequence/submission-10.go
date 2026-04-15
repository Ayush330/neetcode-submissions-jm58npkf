func longestConsecutive(nums []int) int {
    umap := make(map[int]struct{})
    for _, num := range nums{
        umap[num] = struct{}{}
    }

    // start the map sequence
    answer := 0
    maxm := answer
    for key, _:= range umap{
        start := key-1
        if _, exists := umap[start]; exists{
            continue
        }
        answer = 1
        i := key +1 
        for {
            if _, exists := umap[i]; exists{
                i++
                answer++
            }else{
                break
            }
        }
        maxm = max(answer, maxm)
        answer = 0
    }
    return maxm
}

func max(a, b int)int{
    if a>b{
        return a 
    }
    return b
}
