func isValidBST(root *TreeNode) bool {
    return solve(root, math.MinInt64, math.MaxInt64)
}

func solve(root *TreeNode, min, max int) bool {
    if root == nil {
        return true
    }

    // Node value must be strictly within (min, max)
    if root.Val <= min || root.Val >= max {
        return false
    }

    // Left subtree: max bound becomes current node's value
    // Right subtree: min bound becomes current node's value
    return solve(root.Left, min, root.Val) && solve(root.Right, root.Val, max)
}

