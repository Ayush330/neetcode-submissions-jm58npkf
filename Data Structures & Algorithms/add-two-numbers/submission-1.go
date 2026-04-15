/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer := &ListNode{Val : -1}
	result := answer
	carry := 0
	for l1 != nil || l2 != nil{
		v1 := 0
		v2 := 0
		if l1 != nil{
			v1 = l1.Val
		}
		if l2 != nil{
			v2 = l2.Val
		}
		sum := v1 + v2 + carry
		digit := sum%10
		carry = sum/10
		answer.Next = &ListNode{Val : digit}
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
		answer = answer.Next
	}

	if carry > 0{
		answer.Next = &ListNode{Val : carry}
	}

	return result.Next
}
