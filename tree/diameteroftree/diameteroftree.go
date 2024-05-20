package diameteroftree

import tree "algo/pkg"

// 二叉树的最大直径，指树中任意两个结点之间的最大路径长度（两个结点路径中‘——’的个数）；
// 1.求当前结点的左右子树的最大深度lenL, lenR(左右子树的最长链路的结点个数)
// 2.因此，当前结点的最大直径为 lenL+lenR-2
// 时间复杂度：O(N)，需要遍历树中的每一个结点
// 空间复杂度：O(Height)
func diameterOfBinaryTree(root *tree.TreeNode) int {
	//初始化一个ans用于保存结果
	ans := 0
	//声明一个dfs深度优先搜索函数，入参为根节点，出参为当前结点的最大深度（即是当前结点左右子树的最大结点个数）
	var dfs func(node *tree.TreeNode) int
	dfs = func(node *tree.TreeNode) int {
		//递归终止条件
		if node == nil {
			return 0
		}
		//左子树的最大深度
		lenL := dfs(node.Left)
		lenR := dfs(node.Right)
		ans = max(ans, lenL+lenR)
		return max(lenL, lenR) + 1
	}
	dfs(root)
	return ans
}
