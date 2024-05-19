package pkg

// 定义链表的结点（单向）
type ListNode struct {
	Val  int
	Next *ListNode
}

// 初始化链表结点
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}
