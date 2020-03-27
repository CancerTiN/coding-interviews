package main

import "fmt"

type binaryTreeNode struct {
	Value int
	Left  *binaryTreeNode
	Right *binaryTreeNode
}

func (root *binaryTreeNode) preorderTraversal() {
	if root == nil {
		return
	}
	fmt.Printf("%d\t", root.Value)
	if root.Left != nil {
		root.Left.preorderTraversal()
	}
	if root.Right != nil {
		root.Right.preorderTraversal()
	}
}

func (root *binaryTreeNode) inorderTraversal() {
	if root == nil {
		return
	}
	if root.Left != nil {
		root.Left.inorderTraversal()
	}
	fmt.Printf("%d\t", root.Value)
	if root.Right != nil {
		root.Right.inorderTraversal()
	}
}

func construct(preorderSlice []int, inorderSlice []int) (root *binaryTreeNode) {
	if len(preorderSlice) == 0 || len(inorderSlice) == 0 {
		return
	}
	v, i := preorderSlice[0], 0
	for ; inorderSlice[i] != v; i++ {
	}
	root = &binaryTreeNode{Value:v}
	root.Left = construct(preorderSlice[1:i+1], inorderSlice[:i])
	root.Right = construct(preorderSlice[i+1:], inorderSlice[i+1:])
	return
}

func main() {
	preorderSlice := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorderSlice := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root := construct(preorderSlice, inorderSlice)
	root.preorderTraversal()
	fmt.Println()
	root.inorderTraversal()
}
