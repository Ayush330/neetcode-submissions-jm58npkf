/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
	count = solve(root, count, root.Val)
	fmt.Println("count is: ", root.Val)
	return count
}



func solve(root *TreeNode, count int, curr int)int{
	if root == nil{
		return 0
	}
	currVal := root.Val
	if currVal >= curr{
		return 1 + solve(root.Left, count+1, currVal) + solve(root.Right, count+1, currVal)
	}
	return 0 + solve(root.Left, count, curr) + solve(root.Right, count, curr)
}
