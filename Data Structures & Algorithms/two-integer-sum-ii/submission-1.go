func twoSum(numbers []int, target int) []int {
    i := 0; j := len(numbers)-1
    for i<j{
        if numbers[i]+numbers[j] == target{
            break
        }
        if numbers[i]+numbers[j] > target{
            j--
        }else{
            i++
        }
    }

    if i == j{
        return []int{-1, -1}
    }
    return []int{i+1, j+1}
}
