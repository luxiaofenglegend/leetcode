package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := inner113(root, sum)
	for i, _ := range res {
		res[i] = reverse113(res[i])
	}
	return res
}

func inner113(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	// if leaf
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			tmp := []int{sum}
			res := make([][]int, 0)
			res = append(res, tmp)
			return res
		} else {
			return nil
		}
	}
	// if not leaf
	sum -= root.Val
	leftPath := inner113(root.Left, sum)
	rightPath := inner113(root.Right, sum)
	if leftPath != nil {
		for idx, _ := range leftPath {
			leftPath[idx] = append(leftPath[idx], root.Val)
			//leftPath[idx] = reverse113(leftPath[idx])
		}
	}
	if rightPath != nil {
		for idx, _ := range rightPath {
			rightPath[idx] = append(rightPath[idx], root.Val)
			//rightPath[idx] = reverse113(rightPath[idx])
		}
	}
	return append(leftPath, rightPath...)
}

func reverse113(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}
	i, j := 0, len(arr)-1
	for j-i > 0 {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		i++
		j--
	}
	return arr
}
