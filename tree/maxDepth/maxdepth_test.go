package maxDepth

import (
	tree "algo/pkg"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	n1 := tree.NewTreeNode(3)
	n2 := tree.NewTreeNode(9)
	n3 := tree.NewTreeNode(20)
	n4 := tree.NewTreeNode(15)
	n5 := tree.NewTreeNode(7)
	//链接成二叉树
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	//调用递归实现的最大深度
	res1 := maxDepth_rec(n1)
	fmt.Println(res1)

	res2 := maxDepth_que(n1)
	fmt.Println(res2)
}
