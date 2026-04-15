import "slices"

func topKFrequent(nums []int, k int) []int {
    type pair struct{
        num int
        count int
    }
    mp := make(map[int]pair)
    for _, num := range nums{
        p, exists := mp[num]
        if exists{
            p.count++
            mp[num] = p

        }else{
            mp[num] = pair{num, 1}
        }
    }

    lst := []pair{}
    for _, value := range mp{
        lst = append(lst, value)
    }
    slices.SortFunc(lst, func(a, b pair)int{
        return b.count - a.count
    })

    answer := make([]int, k)
    for i:=0; i<k; i++{
        answer[i] = lst[i].num
    }
    return answer
}
