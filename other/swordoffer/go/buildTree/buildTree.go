package buildtree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	node := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	node.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return node
}
