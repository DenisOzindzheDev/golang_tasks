package main

// https://leetcode.com/problems/find-bottom-left-tree-value/description/?envType=daily-question&envId=2024-02-28 TODO
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func findBottomLeftValue(root *TreeNode) int {
	return 0
}

func search(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		// we reached leaf node
		return root.Val
	}

	res := root.Left.Val

	return res
}

// func binarySearch(root *TreeNode) []int {
// 	//var res []int
// 	if root == nil {
// 		return []int{}
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return []int{root.Val}
// 	}
// 	return append(binarySearch(root.Left), binarySearch(root.Right)...)
// }

/*
	пройтись по всему циклу и проверять левые деревья последними и в последней иттерации вернуть значение у дерева

*/
