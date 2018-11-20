/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode)
    carry := 0
    previous := head
    for l1 != nil || l2 != nil || carry != 0 {
        current := new(ListNode)
        previous.Next = current
        
        digit1 := 0
        digit2 := 0
        if l1 != nil {
            digit1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            digit2 = l2.Val
            l2 = l2.Next
        }
        
        sum := digit1 + digit2 + carry
        if sum >= 10 {
            carry = 1
            sum -= 10
        } else {
            carry = 0
        }
        current.Val = sum
        previous = current
    }
    return head.Next
}