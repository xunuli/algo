package maxDepth

import tree "algo/pkg"

// 深度优先搜索，递归实现
// 思路：当前根节点的最大深度 = max(左子树的最大深度， 右子树的最大深度)+1
// 时间复杂度：O(n)，因为要遍历一遍所有结点
// 空间复杂度：O(height)，二叉树的高度
func maxDepth_rec(root *tree.TreeNode) int {
	//递归终止条件
	if root == nil {
		//表示深度为0
		return 0
	}
	//左右子树的最大深度
	ldepth := maxDepth_rec(root.Left)
	rdepth := maxDepth_rec(root.Right)
	//返回结果
	return max(ldepth, rdepth) + 1
}

//广度优先搜索：层序遍历，一层一层的从左到右依次遍历，采用队列实现
//思路：每遍历一层，层高加1

//广度优先搜索实现思路：采用队列实现+双层循环
//1.首先将根节点加入队列中，此时队列的长度大于0，为len1
//2.内层循环按照len1大小去遍历队列的每一个结点（内存循环表示的是某一层次的结点）
//3.内层循环将遍历到的结点依次出队，并将其左右非空子节点入队
//4.当达到len1长度时，表示当层结点已经遍历完成，此时更新len1继续遍历下一层
//时间复杂度：O(n)，每个结点都会访问到

func maxDepth_que(root *tree.TreeNode) int {
	//当根节点为nil，说明层数为0
	if root == nil {
		return 0
	}
	//结果
	ans := 0
	//显示的维护一个队列
	queue := []*tree.TreeNode{}
	//将根节点入对
	queue = append(queue, root)
	//开循环去遍历每一层的每一个结点
	for len(queue) > 0 {
		//表示每一层结点的长度
		len_ceng := len(queue)
		for len_ceng > 0 {
			//将队列中的结点出栈
			node := queue[0]
			queue = queue[1:]
			//将node的左右儿子入栈
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			len_ceng--
		}
		//遍历完一层，ans+1
		ans++
	}
	return ans
}
