package main

import "reflect"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(binarySearch(root1), binarySearch(root2))
}

func binarySearch(root *TreeNode) []int {
	//var res []int
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(binarySearch(root.Left), binarySearch(root.Right)...)
}

/*

SOLUTION:

Считается что мы добавляем эллемент в срез в случае если у него нету рутов по бокам
Мы идем по деревеву рекурсивно если нету слева то бьем справа - если нету слева то справа и так пока
не найдем нужный эллемент.

*/
