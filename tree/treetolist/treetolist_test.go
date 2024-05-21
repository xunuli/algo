package treetolist

import (
	tree "algo/pkg"
	"testing"
)

func Test(t *testing.T) {
	//调用pkg中的构造树结点函数，初始化树结点
	n1 := tree.NewTreeNode(1)
	n2 := tree.NewTreeNode(2)
	n3 := tree.NewTreeNode(3)
	n4 := tree.NewTreeNode(4)
	n5 := tree.NewTreeNode(5)
	n6 := tree.NewTreeNode(6)
	//将节点链接成树
	n1.Left = n2
	n1.Right = n4
	n2.Right = n3
	n4.Left = n5
	n4.Right = n6
	//调用函数
	flatten(n1)
	tree.PrintTree(n1)
}
