/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    array := make([]int, 0)
	array = solve(root, array)
	fmt.Println("The returned slice is: ", array)
	return array[k-1]
}


func solve(root *TreeNode, slice []int) []int{
	if root == nil{
		return slice
	}
	if root.Left != nil{
		slice = solve(root.Left, slice)
	}

	slice = append(slice, root.Val)

	if root.Right != nil{
		slice = solve(root.Right, slice)
	}

	return slice

}
