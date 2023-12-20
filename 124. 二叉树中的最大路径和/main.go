package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var ans = math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := max(dfs(node.Left), 0)
		rightSum := max(dfs(node.Right), 0)

		sum := leftSum + rightSum + node.Val
		ans = max(ans, sum)

		return max(leftSum, rightSum) + node.Val
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
