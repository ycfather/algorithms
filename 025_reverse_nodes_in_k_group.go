package main

import (
	"fmt"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}

	if size < k {
		return head
	}

	var nHead *ListNode = nil
	var cursor *ListNode = nil
	var preSliceTail *ListNode = nil
	var curSliceTail *ListNode = nil
	p := head
	for size >= k {
		cursor = p
		var prev *ListNode = nil
		for i := 0; i < k; i++ {
			if i == 0 {
				curSliceTail = cursor
			}
			next := cursor.Next
			cursor.Next = prev
			prev = cursor
			cursor = next
		}

		p = cursor

		if preSliceTail == nil {
			preSliceTail = head
		} else {
			preSliceTail.Next = prev
			preSliceTail = curSliceTail
		}

		if nHead == nil {
			nHead = prev
		}

		size -= k
	}

	if cursor != nil && preSliceTail != nil {
		preSliceTail.Next = cursor
	}

	return nHead
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
	head.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}

	fmt.Printf("original : %s, after swapped : %s\n", PrettyFormat(head), PrettyFormat(reverseKGroup(head, 2)))
}
