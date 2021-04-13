/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	stackLen := 0
	stack := make([]*TreeNode, 0)
	p_current := root
	if p_current == nil {
		return
	}
	for {
		if p_current.Right != nil && p_current.Left != nil {
			stack = append(stack, p_current.Right)
			stackLen++
		}
		if p_current.Left != nil {
			p_current.Right = p_current.Left
			p_current.Left = nil
		}
		if p_current.Right == nil {
			if stackLen == 0 {
				return
			}
			p_current.Right = stack[stackLen-1]
			stackLen--
			stack = stack[:stackLen]
		}
		p_current = p_current.Right
	}
}

// @lc code=end

