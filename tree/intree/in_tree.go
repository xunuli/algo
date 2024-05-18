package intree

import (
	. "algo/pkg"
)

// 递归实现中序遍历：先左--中根（打印）--后右
// 递归终止条件：当节点为nil时终止，逐层向上返回
// 1.先递归搜索左子树
// 2.保存根结点
// 3.后递归搜索右子树
// 时间复杂度：O(n)，二叉树中的每个结点都会被访问一次
// 空间复杂度：O(n)，取决于栈的深度，最坏情况下为一条链为n
func inorderTraversal_res(root *TreeNode) []int {
	//存储结果的数组
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		//递归终止条件
		if node == nil {
			return
		}
		//先左
		dfs(node.Left)
		//中根，保存
		res = append(res, node.Val)
		//后右
		dfs(node.Right)
	}
	//调用递归函数dfs
	dfs(root)
	return res
}

//中序遍历：先左--中根（保存）--后右
//迭代实现：采用栈实现；
//1.当node当前结点不为nil时入栈，node = node.Left，一直循环到叶子结点
//2.当无法入栈时，栈顶元素可以看作当前的根结点，保存栈顶元素，并出栈，node = node.Right
//3.此时node如果不为nil，则继续入栈，node = node.Left
//4.当前node为nil，栈顶元素作为根节点出栈，保存栈顶元素。。。

func inorderTraversal_stk(root *TreeNode) []int {
	//保存结果的数组
	res := []int{}
	//显示的维护一个栈，go语言中采用动态数组实现
	stk := []*TreeNode{}
	//迭代进栈，先一直左进栈，出栈顶元素，右进栈，循环（把每一个结点都看作根节点处理）
	node := root
	for node != nil || len(stk) > 0 { //当node==nil时需要出栈，只要栈中有元素就继续迭代
		//根节点一直向左循环进栈
		for node != nil {
			stk = append(stk, node)
			node = node.Left
		}
		//当前结点为nil时，需要当前根结点的左儿子为nil，需要将栈顶元素出栈作为根节点保存
		node = stk[len(stk)-1]
		res = append(res, node.Val)
		stk = stk[:len(stk)-1] //出栈
		//当前结点指向右儿子
		node = node.Right
	}
	return res
}
