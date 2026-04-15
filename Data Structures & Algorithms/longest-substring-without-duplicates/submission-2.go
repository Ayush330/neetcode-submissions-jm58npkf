func lengthOfLongestSubstring(s string) int {
	return solve(s)
}

func solve(s string) int{
    if len(s) <= 1{
        return len(s)
    }
    
    mp := make(map[byte]int)
    start := 0
    curr := start
    ans := 0
    count := 0
    indices := [2]int{}
    for curr < len(s){
        index, ok := mp[s[curr]]
        if !ok{
            mp[s[curr]] = curr
            curr++
            count++
        } else if index < start{
            mp[s[curr]] = curr
            curr++
            count++
        }else{  
            if count > ans{
                ans = count
                indices[0] = start
                indices[1] = curr-1
            }
            count = curr - index - 1
            start = index+1
        }
    }
    if count > ans{
                ans = count
                indices[0] = start
                indices[1] = curr-1
            }
    return indices[1] - indices[0] + 1
 }
 