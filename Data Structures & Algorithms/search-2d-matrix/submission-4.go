func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])

    s := m*n

    start := 0; end := s-1
    for start <= end{
        mid := start + (end-start)/2
        row := mid/n
        col := mid%n
        if matrix[row][col] == target{
            return true
        }else if matrix[row][col] > target{
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return false
}
