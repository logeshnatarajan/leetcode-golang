/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
    var prev *ListNode
	for head!= nil {
        nextnode:=head.Next
        head.Next=prev
        prev=head
        head=nextnode
        
	}

	return prev
}
