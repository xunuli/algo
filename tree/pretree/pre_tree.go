package pretree

import . "algo/pkg"

//前序遍历递归实现：先根（保存结果）--中左--后右
//递归终止条件node==nil，逐层向上返回
//1.保存根节点（结果）
//2.递归搜索左子树
//3.递归搜索右子树
// 时间复杂度：O(n)，二叉树中的每个结点都会被访问一次
// 空间复杂度：O(n)，取决于栈的深度，最坏情况下为一条链为n

func preorderTraversal_res(root *TreeNode) []int {
	//保存结果的数组
	res := []int{}
	//声明递归搜索函数dfs
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		//递归终止条件
		if node == nil {
			return
		}
		//先保存根节点
		res = append(res, node.Val)
		//递归搜索左结点
		dfs(node.Left)
		//递归搜索右子树
		dfs(node.Right)
	}
	//调用dfs函数从root开始递归
	dfs(root)
	return res
}

// 迭代实现：采用栈实现
// 中序遍历：先根（保存）--中左--后右
// 迭代实现：采用栈实现；
// 1.当node当前结点（看作根）不为nil时入栈，并保存结果；
// 2.node = node.Left，入栈（看作根节点），保存结果，当node==nil时出栈，此时node为栈顶元素
// 3.node = node.Right，入栈，保存结果，node == nil 出栈.循环。。。
func preorderTraversal_stk(root *TreeNode) []int {
	//定义一个数组用于保存结果
	res := []int{}
	//显示的维护一个栈
	stk := []*TreeNode{}
	node := root //当前结点，看作根节点
	//循环遍历
	for node != nil || len(stk) > 0 {
		//先根中左，将根节点入栈，并保存结果，然后指向其左儿子，循环一直入栈
		for node != nil {
			stk = append(stk, node)
			res = append(res, node.Val)
			node = node.Left
		}
		//当前结点的左儿子为nil时，无法继续向左入栈，将其当作根节点，出栈，并指向其有儿子
		node = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		node = node.Right
	}
	return res
}
