func findDuplicate(nums []int) int {
    umap := make(map[int]struct{})
    for _, num := range nums{
        if _, ok := umap[num]; ok == true{
            return num
        }
        umap[num] = struct{}{}
    }
    return -1
}
