/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var hd *ListNode
    curr := hd
    var node *ListNode
    for l1 != nil || l2 != nil{
        if l1 != nil && l2 != nil{
            sum := l1.Val + l2.Val + carry
            carry = sum / 10
            digit := sum % 10
            node = &ListNode{digit, nil}
            l1 = l1.Next
            l2 = l2.Next
        }else if l1 == nil{
            sum := l2.Val + carry
            carry = sum / 10
            digit := sum % 10
            node = &ListNode{digit, nil}
            l2 = l2.Next
        }else{
            sum := l1.Val + carry
            carry = sum / 10
            digit := sum % 10
            node = &ListNode{digit, nil}
            l1 = l1.Next
        }
        if hd == nil{
            curr = node
            hd = node
        }else{
            curr.Next = node
            curr = curr.Next
        }
    }
    //fmt.Println("Node is: ", node)
    if carry > 0{
        node := &ListNode{carry, nil}
        curr.Next = node
    }

    return hd

}
