package main

import (
	"algo/neet"
	"fmt"
)

func main() {
	//test_contest_ozon.Stickers()
	//test_contest_ozon.Notifications()
	//test_contest_ozon.Championship()
	//test_contest_ozon.Virus()
	root := &neet.TreeNode{
		Val:  1,
		Left: nil,
		Right: &neet.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	ans := neet.InvertTree(root)
	fmt.Println(ans, ans.Right)
}
