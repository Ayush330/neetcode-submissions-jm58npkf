func goodNodes(root *TreeNode) int {
    // Pass the root's own value as the initial maximum to beat.
    // This ensures the root is always counted as the first "good" node.
    return solve(root, root.Val)
}

func solve(node *TreeNode, maxSoFar int) int {
    if node == nil {
        return 0
    }

    // Is this node good?
    isGood := 0
    newMax := maxSoFar

    if node.Val >= maxSoFar {
        isGood = 1
        // ONLY update the maximum if the current node is larger
        newMax = node.Val
    }

    // Recursively count good nodes in branches
    // Note: If node.Val was smaller than maxSoFar, 
    // we pass the old maxSoFar (newMax) down.
    return isGood + solve(node.Left, newMax) + solve(node.Right, newMax)
}