package June_LeetCoding_Challenge

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for n := len(stack); n != 0; n = len(stack) {
		node := stack[n-1]
		stack = stack[:n-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
