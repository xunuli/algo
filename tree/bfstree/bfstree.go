package bfstree

import (
	tree "algo/pkg"
)

// BFS采用队列实现，二叉树的层序遍历，一层一层的依次遍历，
// 双层循环
// 1.第一层循环用于遍历树的总高度
// 2.第二层循环用于遍历每一层的所有结点，遍历节点时，先将该结点出队，在将该节点的左右子结点入队
// 第一层循环条件：len(queue)>0
// 第二层循环条件：len_ceng=len(queue)>0
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	//初始化一个res用于保存结果
	res := [][]int{}
	//显示维护一个队列
	queue := []*tree.TreeNode{}
	//先将根节点入队列
	queue = append(queue, root)
	//第一层循环次数为树的高度
	for len(queue) > 0 {
		//第二层循环次数为每一层的结点个数
		len_ceng := len(queue)
		vals := []int{}
		for len_ceng > 0 {
			//将遍历到的结点出队列，将其的左右子结点如队列
			node := queue[0]
			queue = queue[1:]
			vals = append(vals, node.Val)
			len_ceng--
			//将当前结点的左右子节点加入到队列中
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, vals)
	}
	return res
}
