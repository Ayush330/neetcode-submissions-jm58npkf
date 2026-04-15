func exist(board [][]byte, word string) bool {
    for i:=0; i<len(board); i++{
        for j:=0; j<len(board[0]); j++{
            if board[i][j] == word[0]{
                res := helper(i, j, 0,board, word)
                if res{
                    return true
                }
            }
        }
    }
    return false
}

func helper(start, end, curr int, board [][]byte, word string)bool{
    rowEnd := len(board)
    colEnd := len(board[0])
    if curr == len(word){
        return true
    }
    if start >= rowEnd || start < 0 || end >= colEnd || end < 0{
        return false
    }
    
    if board[start][end] != word[curr]{
        return false
    }
    tmp := board[start][end]
    board[start][end] = '#'
    res := helper(start+1, end, curr+1, board, word) ||
            helper(start, end+1, curr+1, board, word) ||
            helper(start-1, end, curr+1, board, word) ||
            helper(start, end-1, curr+1, board, word)
    board[start][end] = tmp
    return res
}
