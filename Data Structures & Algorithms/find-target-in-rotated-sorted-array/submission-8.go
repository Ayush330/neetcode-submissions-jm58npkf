func search(nums []int, target int) int {
    start := 0
    end := len(nums)-1

    for start < end{
        mid := start + (end-start)/2
        if nums[mid] > nums[end]{
            start = mid+1
        }else{
            end = mid
        }
    }

    index := -1
    x := 0
    y := start-1
    for x<=y{
        mid := x + (y-x)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            x = mid+1
        }else{
            y = mid-1
        }
    }

    x = start 
    y = len(nums)-1
    for x <= y{
         mid := x + (y-x)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            x = mid+1
        }else{
            y = mid-1
        }
    }
    return index
}
