package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs(非递归写法)
type pair struct {
	node  *TreeNode
	depth int
}

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	maxHigh := -1
	stack := []*pair{{root, 0}}

	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := temp.node
		if temp.depth > maxHigh {
			maxHigh = temp.depth
			ans = append(ans, node.Val)
		}

		if node.Left != nil {
			stack = append(stack, &pair{node.Left, temp.depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, &pair{node.Right, temp.depth + 1})
		}
	}

	return ans
}

// dfs(递归写法)
//func rightSideView(root *TreeNode) []int {
//	var ans []int
//	dfs(&ans, root, 0)
//	return ans
//}
//
//func dfs(arr *[]int, node *TreeNode, depth int) {
//	if node == nil {
//		return
//	}
//
//	if len(*arr) == depth {
//		*arr = append(*arr, node.Val)
//	}
//	dfs(arr, node.Right, depth+1)
//	dfs(arr, node.Left, depth+1)
//}

// bfs
//func rightSideView(root *TreeNode) []int {
//	var ans []int
//	if root == nil {
//		return ans
//	}
//
//	queue := []*TreeNode{root}
//	for len(queue) > 0 {
//		cLen := len(queue)
//		for i := 1; i > cLen; i++ {
//			temp := queue[0]
//			queue = queue[1:]
//
//			if i == cLen {
//				ans = append(ans, temp.Val)
//			}
//			if temp.Left != nil {
//				queue = append(queue, temp.Left)
//			}
//			if temp.Right != nil {
//				queue = append(queue, temp.Right)
//			}
//		}
//	}
//	return ans
//}
