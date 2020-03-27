package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListTable(head *ListNode, size int) {
	rand.Intn(size)
	temp := head
	for i := 0; i < size; i++ {
		temp.Next = &ListNode{Val: rand.Intn(size)}
		temp = temp.Next
	}
}

func printList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	for temp := head.Next; temp != nil; temp = temp.Next {
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}

func printListReverselyIteratively(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slice := make([]int, 0)
	for temp := head.Next; temp != nil; temp = temp.Next {
		slice = append(slice, temp.Val)
	}
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Printf("%d\t", slice[i])
	}
	fmt.Println()
}

func printListReverselyRecursively(node *ListNode) {
	if node == nil {
		return
	}
	printListReverselyRecursively(node.Next)
	fmt.Printf("%d\t", node.Val)
}

func main() {
	head := &ListNode{}
	createListTable(head, 9)
	printList(head)
	printListReverselyIteratively(head)
	printListReverselyRecursively(head.Next)
}
