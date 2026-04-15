func wordBreak(s string, wordDict []string) bool{
    mp := make(map[string]bool)
    return wordBreakHelper(s, wordDict, mp)
}

func wordBreakHelper(s string, wordDict []string, mp map[string]bool) bool {
    if len(s) == 0{
        mp[s] = true
        return true
    }
    if val, exists := mp[s]; exists{
        return val
    }
    for i:=0; i<len(s); i++{
        ans := wordBreakHelper(s[i+1:], wordDict, mp)
        mp[s[i+1:]] = ans
        if isMember(s[:i+1], wordDict) && ans{
            return true
        }
    }    
    return false
}

func isMember(s string, dict []string)bool{
    for _, str := range dict{
        if s == str{
            return true
        }
    }
    return false
}
