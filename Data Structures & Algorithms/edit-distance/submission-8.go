func minDistance(word1 string, word2 string) int {
    return solveDP(word1, word2)
}


// we will make a dp state
// i, j will represent the indexs or lenths, basicaly the substtrings and will find a solution to that


func solveDP(word1, word2 string)int{
    l1 := len(word1)
    l2 := len(word2)   

    dp := make([][]int, l1+1)
    for i:=0; i<=l1; i++{
        dp[i] = make([]int, l2+1)
    } 


    // initialize the base cases


    for i:=0; i<=l1; i++{
        dp[i][0] = i
    }

    for i:=0; i<=l2; i++{
        dp[0][i] = i
    }

    for i:=1; i<len(dp); i++{
        for j:=1; j<len(dp[0]); j++{
            if word1[i-1] == word2[j-1]{
                dp[i][j] = dp[i-1][j-1]
            }else{
                // insert
                insert := dp[i][j-1]
                // replace 
                replace := dp[i-1][j-1]
                // delete
                del := dp[i-1][j]
                dp[i][j] = 1 + min(insert, min(replace, del))
            }
        }
    }

    return dp[len(dp)-1][len(dp[0])-1]
}

func solve(i, j int, word1, word2 string) int{
    l1 := len(word1)
    l2 := len(word2)

    if i == l1{
        return max(abs(l2-j), 0)
    }

    if j == l2{
        return max(abs(l1-i), 0)
    }
    if word1[i] == word2[j]{
        return solve(i+1, j+1, word1, word2)
    }
    insert := solve(i, j+1, word1, word2)
    del := solve(i+1, j, word1, word2)
    replace := solve(i+1, j+1, word1, word2)

    return 1 + min(insert, min(del, replace))
}

func abs(i int)int{
    if i > 0{
        return i
    }
    return -1*i
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}