package sortarrtobst

import tree "algo/pkg"

//题目：给定一个整数数组，升序排序，将其转换为平衡二叉搜索树

//前置概念：
//	1、二叉搜索树：
//			1.所有结点都满足：左子树的所有结点值 < 根结点的值 < 右子树的所有结点值；
//			2.无重复结点；
//	2、二叉搜索树的中序遍历是升序排列数组；
//	3、平衡二叉搜索树中的平衡：所有结点的左子树和右子树的深度只差的绝对值不超过1；

//将升序排列数组转化为平衡二叉搜索树：分治递归实现
//	1、升序排列数组==》是二叉树的升序排列；
//	2、要保证二叉搜索树的平衡性，可以采用将中间节点，或中间结点左右相邻的值作为根节点，对左右两侧分割为左右子树
//	3、左右两侧的树仍然需要满足是平衡二叉搜索树，因此可以采用同样的方式对左右两侧递归实现

func sortedArrayToBst(nums []int) *tree.TreeNode {
	//递归终止条件：当数组的长度为0时，返回nil
	if len(nums) == 0 {
		return nil
	}
	//找中间结点的下标
	mid := len(nums) / 2
	//确定根节点
	root := &tree.TreeNode{Val: nums[mid]}
	//根节点的左节点仍然是平衡二叉搜索树
	root.Left = sortedArrayToBst(nums[:mid]) //左闭右开
	root.Right = sortedArrayToBst(nums[mid+1:])
	//返回根节点
	return root
}
