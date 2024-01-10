package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	var indexes = make(map[int]int)
	for i := 0; i < n; i++ {
		indexes[inorder[i]] = i
	}

	var dfs func(preLeft, preRight, inLeft, inRight int) *TreeNode
	dfs = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight {
			return nil
		}
		preRoot := preLeft
		rootVal := preorder[preRoot]

		inRoot := indexes[rootVal]
		root := &TreeNode{Val: rootVal}
		leftSubTreeSize := inRoot - inLeft

		root.Left = dfs(preLeft+1, preLeft+leftSubTreeSize, inLeft, inRoot-1)
		root.Right = dfs(preLeft+leftSubTreeSize+1, preRight, inRoot+1, inRight)
		return root
	}

	return dfs(0, n-1, 0, n-1)
}
