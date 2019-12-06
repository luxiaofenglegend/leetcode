package tree

/*
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]

Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	length := len(preorder)
	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	if len(preorder) == 1 {
		return root
	}

	idx := 0
	for idx < length && inorder[idx] != rootVal {
		idx++
	}
	lfLength := idx
	//rgtLength := length - lfLength - 1
	root.Left = buildTree(preorder[1:1+lfLength], inorder[0:lfLength])
	root.Right = buildTree(preorder[lfLength+1:], inorder[idx+1:])
	return root
}
