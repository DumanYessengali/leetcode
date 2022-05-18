package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data > data {
		child := BTreeInsertData(root.Left, data)
		root.Left = child
		child.Parent = root
	} else {
		child := BTreeInsertData(root.Right, data)
		root.Right = child
		child.Parent = root
	}
	return root
}
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
}
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if BTreeSearchItem(root.Left, elem) != nil {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MaxHeight(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right)) + 1
}
func MaxHeight(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right) {
		return true
	}
	return false
}
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		ElemPrint(root, i, f)
	}
}
func ElemPrint(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else {
		ElemPrint(root.Left, i-1, f)
		ElemPrint(root.Right, i-1, f)
	}
}
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)
}
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left == node {
		root.Left = rplc
		root.Left.Parent = root
	} else if root.Right != nil && root.Right == node {
		root.Right = rplc
		root.Right.Parent = root
	} else {
		BTreeTransplant(root.Left, node, rplc)
		BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data > node.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if root.Data < node.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		root.Data = BTreeMin(root.Right).Data
		root.Right = BTreeDeleteNode(root.Right, node)
	}
	return root
}
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}
