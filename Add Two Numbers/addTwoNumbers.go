/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{0,nil}
    var result = head
    carry :=0
    for (l1 != nil || l2 != nil ){
        sum:=0
        if l1 != nil{
            sum+=l1.Val
            l1=l1.Next
        }
        if l2 != nil{
            sum+=l2.Val
            l2=l2.Next
        }
        sum+=carry
        carry=sum/10
        result.Next= &ListNode{sum%10,nil}
        result=result.Next
    }
    if carry!=0{
        result.Next= &ListNode{carry,nil}
        result=result.Next
    }
    return head.Next
}
