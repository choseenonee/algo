package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swapChildrenAndGoDown(treeNode *TreeNode) {
	if treeNode.Right == nil && treeNode.Left == nil {
		return
	}
	if treeNode.Right == nil {
		treeNode.Right = treeNode.Left
		treeNode.Left = nil
		swapChildrenAndGoDown(treeNode.Right)
	} else if treeNode.Left == nil {
		treeNode.Left = treeNode.Right
		treeNode.Right = nil
		swapChildrenAndGoDown(treeNode.Left)
	} else {
		left := treeNode.Left
		treeNode.Left = treeNode.Right
		treeNode.Right = left
		swapChildrenAndGoDown(treeNode.Left)
		swapChildrenAndGoDown(treeNode.Right)
	}
}

func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		swapChildrenAndGoDown(root)
	}

	return root
}
