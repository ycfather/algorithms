package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrettyFormat(head *ListNode) string {

	if head == nil {
		return "nil"
	}

	var str string
	cursor := head
	for cursor != nil {
		str += fmt.Sprintf("[%d]", cursor.Val)
		if cursor.Next != nil {
			str += " ---> "
		}
		cursor = cursor.Next
	}

	return str
}

func Reverse(head *ListNode) *ListNode {
	cursor := head
	var prev *ListNode

	for cursor != nil {
		next := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = next
	}

	return prev
}
