func searchMatrix(matrix [][]int, target int) bool {
    for i:=0; i<len(matrix); i++{
        curr := matrix[i]
        if target <= curr[len(curr)-1]{
            return binarySearch(curr, target)
        }
    }
    return false
}


func binarySearch(matrix []int, target int) bool{
    start := 0
    end := len(matrix)

    for start <= end{
        mid := start + (end-start)/2
        if target == matrix[mid]{
            return true
        }
        if target < matrix[mid]{
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return false
}