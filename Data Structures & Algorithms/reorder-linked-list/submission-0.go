/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // find the middle
	slow := head
	fast := head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := slow.Next
	slow.Next = nil
	slow = reverse(slow)
	secondHalf = reverse(secondHalf)
	firstHalf := head
	for secondHalf != nil {
        // Save our forward paths
        tmp1 := firstHalf.Next
        tmp2 := secondHalf.Next
        
        // The Zig: Left points to Right
        firstHalf.Next = secondHalf
        
        // The Zag: Right points back to the next Left (YOU WERE MISSING THIS!)
        secondHalf.Next = tmp1
        
        // Move our runners forward
        firstHalf = tmp1
        secondHalf = tmp2
    }
}



func reverse(head *ListNode)*ListNode{
	curr := head
	var prev *ListNode
	for curr != nil{
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}