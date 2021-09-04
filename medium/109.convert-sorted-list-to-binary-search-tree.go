/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(h *ListNode) *TreeNode {
	return helper(h, nil)
}

func helper(h, t *ListNode) *TreeNode {
	if h == t {
		return nil
	}

	fast := h
	slow := h
	for fast.Next != t && fast.Next.Next != t {
		fast = fast.Next.Next
		slow = slow.Next
	}

	rs := &TreeNode{Val: slow.Val}
	rs.Left = helper(h, slow)
	rs.Right = helper(slow.Next, t)

	return rs
}
