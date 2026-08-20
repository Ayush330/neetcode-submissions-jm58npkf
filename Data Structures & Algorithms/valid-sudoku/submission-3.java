class Solution {
    public boolean isValidSudoku(char[][] board) {
        int[] row = new int[9];
        int[] col = new int[9];
        int[] grid = new int[9];
        for(int i=0; i<board.length; i++){
            for(int j=0; j<board[i].length; j++){
                if(board[i][j]=='.') continue;
                int n = board[i][j] -'0';
                int mask = 1 << n;
                int g = (i/3)*3 + (j/3);
                if(
                    (mask & row[i]) != 0 ||
                    (mask & col[j]) != 0 ||
                    (mask & grid[g]) != 0
                ) return false;
                row[i] = mask | row[i];
                col[j] = mask | col[j];
                grid[g] = mask | grid[g];
            }
        }
        return true;
    }
}
