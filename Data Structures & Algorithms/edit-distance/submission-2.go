func minDistance(word1 string, word2 string) int {
    mp := make(map[[2]int]int)
    return solve(word1, word2, len(word1), len(word2), mp)
}

func solve(word1, word2 string, i, j int, mp map[[2]int]int)int{
    if val, exists := mp[[2]int{i,j}]; exists{
        return val
    }

    if i == 0{
        return j
    }
    if j == 0{
        return i
    }

    if word1[i-1] == word2[j-1]{
        x := solve(word1, word2, i-1, j-1, mp)
        mp[[2]int{i-1, j-1}] = x
        return x
    }

    a := solve(word1, word2, i-1, j, mp)
    b := solve(word1, word2, i, j-1, mp)
    c := solve(word1, word2, i-1, j-1, mp)
    mp[[2]int{i-1, j}] = a
    mp[[2]int{i, j-1}] = b
    mp[[2]int{i-1, j-1}] = c
    return 1+min(min(a, b), c)
}


func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}
