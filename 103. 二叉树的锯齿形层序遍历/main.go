package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	mark := 1
	for len(queue) > 0 {
		mark = 1 - mark
		curLen := len(queue)
		arr := make([]int, curLen)

		for i := 0; i < curLen; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}

			if mark == 0 {
				arr[i] = temp.Val
			} else {
				arr[curLen-i-1] = temp.Val
			}
		}

		ans = append(ans, arr)
	}

	return ans
}
