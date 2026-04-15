func kthSmallest(root *TreeNode, k int) int {
    count := 0
    result := -1
    
    // We pass the MEMORY ADDRESSES of count and result so the 
    // recursive function can permanently modify them.
    solve(root, &count, &result, k)
    
    return result
}

// Notice count and result are now *int (pointers!)
func solve(root *TreeNode, count *int, result *int, k int) {
    // 1. Base case: Stop if node is nil, or if we already found the answer
    if root == nil || *count >= k {
        return
    }

    // 2. Go Left (Find the smallest numbers first)
    solve(root.Left, count, result, k)

    // 3. Process the current node
    *count++             // We arrived at a node, increase the shared tally!
    if *count == k {     // Is this the magic number?
        *result = root.Val // Save it to the shared result box!
        return             // Pull the fire alarm, we are done!
    }

    // 4. Go Right (Find slightly larger numbers)
    solve(root.Right, count, result, k)
}