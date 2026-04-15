func minDistance(word1 string, word2 string) int {
    return solve(word1, word2, len(word1), len(word2))
}

func solve(word1, word2 string, i, j int)int{
    if i == 0{
        return j
    }
    if j == 0{
        return i
    }

    if word1[i-1] == word2[j-1]{
        return solve(word1, word2, i-1, j-1)
    }

    a := solve(word1, word2, i-1, j)
    b := solve(word1, word2, i, j-1)
    c := solve(word1, word2, i-1, j-1)

    return 1+min(min(a, b), c)
}


func min(a, b int)int{
    if a <b{
        return a
    }
    return b
}
