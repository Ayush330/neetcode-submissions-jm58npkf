func searchMatrix(matrix [][]int, target int) bool {
    return binarySearch(matrix, target)
}


func binarySearch(matrix [][]int, target int) bool{
    start := 0
    end := (len(matrix) * len(matrix[0]))-1
    for start <= end{
        mid := start + (end-start)/2
        if target == matrix[mid/len(matrix[0])][mid%len(matrix[0])]{
            return true
        }
        if target < matrix[mid/len(matrix[0])][mid%len(matrix[0])]{
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return false
}