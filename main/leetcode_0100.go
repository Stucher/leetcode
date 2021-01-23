package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// 当前节点值不等 直接返回false
	if (p == nil && q != nil) || (p != nil && q == nil) || (q.Val != q.Val) {
		return false
	}

	var isSame bool

	// 检查左子树
	if p.Left != nil && q.Left != nil {
		isSame = isSameTree(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		isSame = true
	} else {
		return false
	}

	// 左子树不等 直接返回false
	if !isSame {
		return false
	}

	// 检查右子树
	if p.Right != nil && q.Right != nil {
		isSame = isSameTree(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		isSame = true
	} else {
		return false
	}

	return isSame && q.Val == p.Val
}
