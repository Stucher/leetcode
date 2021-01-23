package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	var notRemain bool

	notRemain = hasPathSum(root.Left, targetSum-root.Val)

	if notRemain {
		return true
	}

	return hasPathSum(root.Right, targetSum-root.Val)
}
