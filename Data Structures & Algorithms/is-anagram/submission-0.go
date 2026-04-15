func isAnagram(s string, t string) bool {
    mp := make(map[rune]int)
    for _, runee := range s{
       mp[runee]++
    }
    for _, runee := range t{
        mp[runee]--
    }
    for _, val := range mp{
        if val != 0{
            return false
        }
    }
    return true
}
