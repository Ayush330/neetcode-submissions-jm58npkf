func groupAnagrams(strs []string) [][]string {
    // iterate over each,
    // create a hash of the key, the value is a list of strings
    mp := make(map[string][]string)
    for _, str := range strs{
        lookup := [26]int{}
        for _, rn := range str{
            index := rn - 'a'
            lookup[index]++
        }
        mp[getKey(lookup)] = append(mp[getKey(lookup)], str)
    }
    res := make([][]string, 0)
    for _, val := range mp{
        res = append(res, val)
    }
    return res
}

func getKey(arr [26]int) string{
    res := ""
    for _, val := range arr{
        res = res + "-" + strconv.Itoa(val)
    }
    return res
}
