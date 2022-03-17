package main

import (
	"fmt"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l, cursor, prev *ListNode
	carry := 0
	cursor1 := l1
	cursor2 := l2
	for cursor1 != nil || cursor2 != nil {
		cursor = &ListNode{Val: 0}
		if cursor1 != nil {
			cursor.Val += cursor1.Val
		}
		if cursor2 != nil {
			cursor.Val += cursor2.Val
		}
		cursor.Val += carry

		carry = cursor.Val / 10
		cursor.Val = cursor.Val % 10

		if l == nil {
			l = cursor
		}

		if prev != nil {
			prev.Next = cursor
		}
		prev = cursor
		if cursor1 != nil {
			cursor1 = cursor1.Next
		}
		if cursor2 != nil {
			cursor2 = cursor2.Next
		}
	}

	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return l
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	carry := 0
	for {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		cur.Val = sum % 10
		carry = sum / 10

		if l1 == nil && l2 == nil {
			break
		}

		cur.Next = new(ListNode)
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head
}

func main() {
	h1 := &ListNode{
		Val: 1,
	}
	h1.Next = &ListNode{
		Val: 8,
	}

	h2 := &ListNode{
		Val: 0,
	}

	fmt.Printf("sum list : %s\n", PrettyFormat(addTwoNumbers2(h1, h2)))
}
