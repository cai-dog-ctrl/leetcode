package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	Search(root)
	var ans = 0
	var PoDU func(root *TreeNode)
	PoDU = func(root *TreeNode) {
		if root == nil {
			return
		}
		var x int
		var y int
		if root.Left == nil {
			x = 0
		} else {
			x = root.Left.Val
		}
		if root.Right == nil {
			y = 0
		} else {
			y = root.Right.Val
		}
		ans += abs(x - y)
		PoDU(root.Left)
		PoDU(root.Right)
	}
	PoDU(root)
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Search(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.Val = root.Val + Search(root.Left) + Search(root.Right)
	return root.Val
}
