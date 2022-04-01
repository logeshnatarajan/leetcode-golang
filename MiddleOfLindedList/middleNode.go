/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    re:=head
    var i,j =0,0
    for head!=nil{
        head=head.Next
        i++
    }
    mid:=(i/2)+1
    for {
        j++
        if j==mid{
            return re
        }
        re=re.Next
    }
}
