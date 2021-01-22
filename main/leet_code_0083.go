package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p, q := head, head.Next

	for q != nil {
		if p.Val != q.Val {
			p.Next = q
			p = q
		}
		q = q.Next
	}
	p.Next = nil

	return head
}
