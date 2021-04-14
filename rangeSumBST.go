package Leetcode

func rangeSumBST(root *TreeNode, low int, high int) int {
    res := 0
    if root == nil {
        return 0
    }
    if (root.Val >= low) && (root.Val <= high) {
        res = res + root.Val
    }
    if root.Val > low {
        res = res + rangeSumBST(root.Left, low, high)
    }
    if root.Val < high {
        res = res + rangeSumBST(root.Right, low, high)
    }
    return res
}