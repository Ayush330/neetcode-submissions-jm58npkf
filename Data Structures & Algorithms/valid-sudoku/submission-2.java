class Solution {
    public boolean isValidSudoku(char[][] board) {
        for(int i=0; i<board.length; i++){
            for(int j=0; j<board[i].length; j++){
                if(board[i][j]!='.'){
                    if(!(row(board, i, j) && col(board, i, j) && grid(board, i, j))) return false;
                }
            }
        }
        return true;
    }

    public boolean row(char[][] board, int x, int y){
        for(int i=0; i<board.length; i++){
            if(i!=x && board[x][y] == board[i][y]) return false;
        }
        return true;
    }

    public boolean col(char[][] board, int x, int y){
        for(int i=0; i<board[0].length; i++){
            if(i!=y && board[x][y] == board[x][i]){
                return false;
            }
        }
        return true;
    }

    public boolean grid(char[][] board, int x, int y){
        int rs = (x/3)*3;
        int cs = (y/3)*3;
        System.out.printf("x:%d\trs:%d\ty:%d\tcs:%d%n", x, rs, y, cs);
        for(int i=rs; i<rs+3; i++){
            for(int j=cs; j<cs+3; j++){
                if(i!=x && j !=y && board[i][j]==board[x][y]) return false;
            }
        }
        return true;
    }
}
