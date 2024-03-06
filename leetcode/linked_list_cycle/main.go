package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	reached := make(map[*ListNode]bool) // if current node reached
	current := head
	for current != nil {
		if reached[current] {
			return true
		}
		reached[current] = true
		current = current.Next
	}
	return false
}

/*
p          |
node 1 4 5 1 5 3
POS  0 1 2 3 4 5

*/
