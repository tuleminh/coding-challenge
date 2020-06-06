package June_LeetCoding_Challenge

func deleteNode(node *ListNode) {
	if node == nil {
		return
	} else {
		if node.Next == nil {
			return
		}
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}
