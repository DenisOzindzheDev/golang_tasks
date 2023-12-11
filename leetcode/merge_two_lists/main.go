package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1.Next == nil {
		return list2
	}
	if list2.Next == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1, list2.Next)
	}
	list2.Next = mergeTwoLists(list1.Next, list2)
	return list2

}
