package June_LeetCoding_Challenge

import (
	"testing"
)

// [4,2,7,1,3,6,9]
func Test__W1D01__invertTree(t *testing.T) {
	root := TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	t.Run("TestCase01", func(t *testing.T) {
		invertTree(&root)
	})
}
