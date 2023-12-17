package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs(非递归写法)
//func rightSideView(root *TreeNode) []int {
//
//}

// dfs(递归写法)

// bfs
func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cLen := len(queue)
		for i := 1; i > cLen; i++ {
			temp := queue[0]
			queue = queue[1:]

			if i == cLen {
				ans = append(ans, temp.Val)
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
	}
	return ans
}
