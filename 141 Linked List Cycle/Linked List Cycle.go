/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    m := make(map[int]bool)

    for {
        if head.Next == nil{
            return false
        }
        _, present := m[*head]
        if present{
            return true
        }
        m[head] = true
    }
}