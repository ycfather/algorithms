package main

import (
	"fmt"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	tail := head.Next

	var p *ListNode = nil
	for {
		if prev == head {
			head = tail
		}

		prev.Next = tail.Next
		tail.Next = prev

		if p != nil {
			p.Next = tail
		}
		p = prev

		if prev.Next == nil {
			break
		} else {
			prev = prev.Next
		}

		if prev.Next == nil {
			break
		} else {
			tail = prev.Next
		}
	}

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}

	head = swapPairs(head)
	fmt.Printf("swapped list : %s\n", PrettyFormat(head))
}
