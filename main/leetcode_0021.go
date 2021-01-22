package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	p, q, r := l1, l2, res

	for p != nil && q != nil {
		if p.Val < q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}

	for p != nil {
		r.Next = p
	}

	for q != nil {
		r.Next = q
	}

	return res.Next
}
