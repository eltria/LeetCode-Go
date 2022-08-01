package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

type Node struct {
	Val  int
	Next *Node
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	isAdd := false

	result := &ListNode{}
	top := result
	for l1 != nil || l2 != nil {
		res := 0

		if l1 != nil {
			res = res + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			res = res + l2.Val
			l2 = l2.Next
		}

		if isAdd {
			res = res + 1
			isAdd = false
		}

		if res >= 10 {
			res = res - 10
			isAdd = true
		}

		result.Next = &ListNode{
			Val:  res,
			Next: nil,
		}
		result = result.Next
	}

	if isAdd {
		result.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
		result = result.Next
	}

	return top.Next
}
