package sortarrtobst

import (
	. "algo/pkg"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//初始化数组
	nums := []int{-10, -3, 0, 5, 9}

	//调用转换函数
	root := sortedArrayToBst(nums)
	//对其进行中序遍历，看是否和原来相等
	res := inordertree(root)
	fmt.Println(res)
}

// 中序遍历：先左--中根--后右
func inordertree(root *TreeNode) []int {
	//保存结果
	res := []int{}
	//定义dfs函数
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		//递归终止条件
		if node == nil {
			return
		}
		//遍历左节点
		dfs(node.Left)
		//打印根节点
		res = append(res, node.Val)
		//遍历右节点
		dfs(node.Right)
	}
	dfs(root)
	return res
}
