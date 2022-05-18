package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	BTreeInsertData(root, 2)
	BTreeInsertData(root, 7)
	BTreeInsertData(root, 1)
	BTreeInsertData(root, 3)
	BTreeApplyInorder(searchBST(root, 2), fmt.Println)
}
