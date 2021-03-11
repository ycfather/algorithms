package main

import (
	"fmt"
)

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Val: -101}
	dummy.Next = head

	pre := dummy
	for head != nil {
		if head.Val >= x {
			h, t := head, head
			for head != nil && head.Val >= x {
				t = head
				head = head.Next
			}
			if head != nil {
				pre.Next = head
				pre = head
				tmp := head.Next
				head.Next = h
				t.Next = tmp
			}
		} else {
			pre = head
			head = head.Next
		}
	}

	return dummy.Next
}

func main() {
	head := Build([]int{1})
	fmt.Printf("origin: %s\n", PrettyFormat(head))
	nHead := partition(head, 2)
	fmt.Printf("partitioned: %s\n", PrettyFormat(nHead))
}
