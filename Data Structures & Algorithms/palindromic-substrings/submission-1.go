func countSubstrings(s string) int {
    mp := make(map[string]bool)
    return countSubstringsH(s, mp)
}

func countSubstringsH(s string, mp map[string]bool) int {
    maxm := 0
    for i:=0; i<len(s); i++{
        for j:=i+1; j<=len(s); j++{
            if val, exists := mp[s[i:j]]; exists{
                if val{
                    maxm++
                }
            }else{
                isP := checkPalindrome(s[i:j])
                //fmt.Println(s[i:j], "  ", isP)
                mp[s[i:j]] = isP
                if isP{
                    maxm++
                }
            }
        }
    }
    return maxm
}

func maxm(a, b int)int{
    if a > b{
        return b
    }
    return a
}



func checkPalindrome(s string)bool{
    start, end := 0, len(s)-1
    for start <= end{
        if s[start] != s[end]{
            return false
        }
        start ++; end--
    }
    return true
}