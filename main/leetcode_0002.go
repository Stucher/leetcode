package main

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overTenFlag := 0
	var res = new(ListNode)
	p, q, r := l1, l2, res
	for p != nil || q != nil {
		pVal := 0
		qVal := 0
		if p != nil {
			pVal = p.Val
			p = p.Next
		}
		if q != nil {
			qVal = q.Val
			q = q.Next
		}
		val := pVal + qVal + overTenFlag
		if val >= 10 {
			overTenFlag = 1
			r.Val = val - 10
		} else {
			overTenFlag = 0
			r.Val = val
		}
		if p != nil || q != nil || overTenFlag != 0 {
			r.Next = new(ListNode)
			r = r.Next
		}
	}
	if overTenFlag != 0 {
		r.Val = 1
	}
	return res
}
