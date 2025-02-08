package buildtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreePreorderInorder(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTreePreorderInorder(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTreePreorderInorder(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func buildTreeInorderPostorder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTreeInorderPostorder(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTreeInorderPostorder(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
	return root
}
