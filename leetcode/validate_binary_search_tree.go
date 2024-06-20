package leetcode

import "math"

func checkSelfValid(lowerLevel, upperLevel, val int) bool {
	return val > lowerLevel && val < upperLevel
}

func isValidBSTMine(root *TreeNode, lowerLevel, upperLevel, count int) bool {
	if root == nil {
		return true
	}

	if count != 0 {
		if !checkSelfValid(lowerLevel, upperLevel, root.Val) {
			return false
		}
	}

	return isValidBSTMine(root.Left, lowerLevel, root.Val, 1) && isValidBSTMine(root.Right, root.Val, upperLevel, 1)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isValidBSTMine(root, -int(math.Pow(2, 31))-1, int(math.Pow(2, 31)), 0)
}
