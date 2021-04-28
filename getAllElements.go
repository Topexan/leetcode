package leetcode

import (
	"sort"
)

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res, arr2 []int
	res = getElements(res, root1)
	arr2 = getElements(arr2, root2)
	for i := range arr2 {
		res = append(res, arr2[i])
	}
	sort.Ints(res)
	return res
}

func getElements(arr []int, root *TreeNode) []int {
	if root == nil {
		return arr
	}
	arr = getElements(arr, root.Left)
	arr = append(arr, root.Val)
	arr = getElements(arr, root.Right)
	return arr
}
