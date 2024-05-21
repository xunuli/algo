package treetolist

import tree "algo/pkg"

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。

//思路：
//常规先序遍历：从根结点开始到左子树的叶子结点不断入栈，入栈时保存结点的信息，只有出栈时将其右儿子加入栈，
//这样如果修改根节点的结构，由于没有保存二叉树的右节点的信息会导致丢失由儿子的信息

//因此采用变种的先序遍历，首先维护一个上一次访问的结点的指针prev，用于转化为链表；
//其次，在先序遍历时先将根结点入栈，然后将其出栈，将其右儿子和左儿子依次入栈，然后栈顶元素出栈，右左儿子入栈一直这样
//因为prev是上一次访问的元素，所以当栈顶元素出栈时curr有，prev.L=nil，prev.R=curr，prev=curr
//迭代条件：for len(stk)>0 {}
//时间复杂度：O(n)
//空间复杂度：O(n)

func flatten(root *tree.TreeNode) {
	//如果root为nil，直接返回
	if root == nil {
		return
	}
	//需要采用迭代实现，显示的维护一个栈，并将根节点入栈
	stk := []*tree.TreeNode{root}
	//维护上一次访问的结点
	prev := &tree.TreeNode{}
	//循环遍历，遍历条件：len(stk)>0
	for len(stk) > 0 {
		//先将栈顶元素出栈作为当前结点（当前根节点）
		curr := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		//如果上一次的访问结点不为nil，则其左儿子为nil，右儿子为curr，形成链表
		if prev != nil {
			prev.Left = nil
			prev.Right = curr
		}
		//移动prev的位置为curr
		prev = curr
		//将当前结点的右，左不为nil的儿子依次入栈顺序不能改变
		if curr.Right != nil {
			stk = append(stk, curr.Right)
		}
		if curr.Left != nil {
			stk = append(stk, curr.Left)
		}
	}
}
