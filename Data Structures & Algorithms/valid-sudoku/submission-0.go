func isValidSudoku(board [][]byte) bool {
    // I can figure out the quadrants
    var row[9][9] bool
    var col[9][9] bool
    var block[9][9] bool

    for i:=0; i<9; i++{
        for j:=0; j<9;j++{
            if board[i][j] == '.'{
                continue
            }
            num := board[i][j] - '1'
            blockIndex := i/3 * 3 + j/3
            if row[num][i] || col[num][j] || block[num][blockIndex]{
                return false
            }
            row[num][i] = true
            col[num][j] = true
            block[num][blockIndex] = true
        }
    }
    return true
}
