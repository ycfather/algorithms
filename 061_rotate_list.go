package main

import (
	"fmt"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	original_tail := head
	n := 1
	for ; original_tail.Next != nil; original_tail = original_tail.Next {
		n++
	}
	original_tail.Next = head

	t := n - k%n - 1
	var new_tail, new_head *ListNode = head, nil
	for i := 0; i < t; i++ {
		new_tail = new_tail.Next
	}
	new_head = new_tail.Next
	new_tail.Next = nil

	return new_head
}

func main() {
	head := Build([]int{1, 2, 3, 4, 5})
	fmt.Printf("rotated: %s\n", PrettyFormat(rotateRight(head, 2)))
}
