package inandprebuildtree

import (
	tree "algo/pkg"
	"testing"
)

func Test(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	//调用函数
	root := buildTree(preorder, inorder)
	tree.PrintTree(root)
}
