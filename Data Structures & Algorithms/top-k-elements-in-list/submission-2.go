func topKFrequent(nums []int, k int) []int {
    fmap := make(map[int]int)

    for _, el := range nums{
        fmap[el]++
    }

    // create buckets

    bucket := make([][]int, len(nums)+1)

    for key, val := range fmap {
        bucket[val] = append(bucket[val], key)
    }

    result := make([]int, 0, k)
    rem := k
    for i:=len(nums); i>=0; i--{
            if len(bucket[i]) <= rem{
                result = append(result, bucket[i]...)
                rem = rem - len(bucket[i])
            }else{
                result = append(result, bucket[i][0:rem]...)
                return result
            }
        
    }

    return result
}
