func longestConsecutive(nums []int) int {
    if len(nums) <= 1{
        return len(nums)
    }
    mp := make(map[int]struct{})
    for _, el := range nums{
        mp[el] = struct{}{}
    }
    max := 1
    for key, _ := range mp{
        if _, ok := mp[key-1]; !ok{ // this is the beginning
            curr := 1
            upKey := key+1
            for {
    if _, ok := mp[upKey]; ok {
        curr++
        max = maxm(max, curr)
        upKey++
    } else {
        break // Stop when the next number is not in the set.
    }
}
        }
    }
    return max
}

func maxm(a, b int) int{
    if a>b{
        return a
    }
    return b
}