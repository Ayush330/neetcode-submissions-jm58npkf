func checkInclusion(s1 string, s2 string) bool {
    l1 := len(s1)
    l2 := len(s2)
    if l1 > l2{
        return false
    }
    var f1 [26]int
    var f2 [26]int

    // initialize the frequency of first one
    for i:=0; i<l1; i++{
        f1[s1[i]-'a']++
        f2[s2[i]-'a']++
    }
    for i:=0; i<l2-l1; i++{
        if f1 == f2{
            return true
        }
        f2[s2[i]-'a']--
        f2[s2[i+l1]-'a']++
    }
    return f1 == f2
}
