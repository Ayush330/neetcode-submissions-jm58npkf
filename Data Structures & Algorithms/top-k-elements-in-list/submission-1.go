func topKFrequent(nums []int, k int) []int {
    mp := make(map[int]int)
    for _, num := range nums{
        mp[num]++
    }

    list := make([][]int, len(nums)+1)
    for i:=0; i<len(nums); i++{
        list[i] = []int{}
    }

    for key, val := range mp{
        list[val] = append(list[val], key)
    }
    //fmt.Println(list)
    ans := make([]int, k)
    curr := 0
    for i:=len(list)-1; i>=0; i--{
        for _, v := range list[i]{
            ans[curr] = v
            curr++
            if curr == k{
                return ans
            }
        }
    }
    return ans
}
