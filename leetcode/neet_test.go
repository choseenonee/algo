package leetcode

import (
	"fmt"
	"testing"
)

func TestValidateBinarySearchTree(t *testing.T) {
	//root := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val:   4,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 6,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	//root := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	root := &TreeNode{
		Val:  -2147483648,
		Left: nil,
		Right: &TreeNode{
			Val:   2147483647,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isValidBST(root))
}
