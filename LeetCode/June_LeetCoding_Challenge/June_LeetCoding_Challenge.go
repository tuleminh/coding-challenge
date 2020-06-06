package June_LeetCoding_Challenge

//
// TreeNode
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
// ListNode
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func NewList(nums []int) *ListNode {
	var (
		head = NewListNode(nums[0])
		node = head
	)
	for i := 1; i < len(nums); i++ {
		node.Next = NewListNode(nums[i])
		node = node.Next
	}
	return head
}
