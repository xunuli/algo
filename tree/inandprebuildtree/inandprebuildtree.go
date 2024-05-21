package inandprebuildtree

import tree "algo/pkg"

//给定两个整数数组 preorder 和 inorder ，
//其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

// 递归实现：思路
// 1.对于前序遍历而言：[根，[左子树的前序遍历]，[右子树的前序遍历]]
// 2.对于中序遍历而言：[[左子树的中序遍历]，根，[右子树的中序遍历]]
// 不管采用什么遍历，左子树和右子树的长度是固定的，因此我们可以分别对某个结点的左右子树进行构造，一直循环
// 时间复杂度：O(n)；
// 空间复杂度：O(h);
func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	//如果前序遍历或中序遍历的结果集合为0，说明树为nil
	if len(preorder) == 0 {
		return nil
	}
	//前序遍历中根节点所在的位置为preorder[0]，构造根节点
	root := &tree.TreeNode{preorder[0], nil, nil}
	//通过找到中序遍历中的根节点位置从而划分左右子树
	i := 0
	for i < len(inorder) {
		if inorder[i] == preorder[0] {
			break
		}
		i++
	}
	//此时中序遍历中根节点的位置为i
	//对于一棵树的根节点而言，不管怎么遍历，左右子树都是固定的
	//通过确定左右子树的长度，可以实现详细递归
	left_len := len(inorder[:i]) //左子树的长度

	//递归实现构造子树
	root.Left = buildTree(preorder[1:left_len+1], inorder[0:i])
	root.Right = buildTree(preorder[left_len+1:], inorder[i+1:])
	return root
}
