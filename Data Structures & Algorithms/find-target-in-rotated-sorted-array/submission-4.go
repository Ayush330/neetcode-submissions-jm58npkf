func search(nums []int, target int) int {
    left := 0; right := len(nums)-1
    for left < right{
        mid := left + (right-left)/2
        if nums[mid] > nums[right]{
            left = mid + 1
        }else{
            right = mid
        }
    }
    fmt.Println(left, right)
    if target >= nums[left] && target <= nums[len(nums)-1]{
        return find(nums, target, left, len(nums)-1)
    }else{
        return find(nums, target, 0, left-1)
    }
}

func find(nums []int, target, left, right int) int{
    for left <= right{
        mid := left + (right-left)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            left = mid+1
        }else{
            right = mid-1
        }
    }
    return -1
}