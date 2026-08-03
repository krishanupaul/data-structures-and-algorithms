// Problem: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergeNode := &ListNode{}
	newList := mergeNode

	for list1 != nil || list2 != nil {
		if list1 == nil {
			mergeNode.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			mergeNode.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			mergeNode.Next = list2
			list2 = list2.Next
		} else {
			mergeNode.Next = list1
			list1 = list1.Next
		}

		mergeNode = mergeNode.Next
	}

	return newList.Next
}
