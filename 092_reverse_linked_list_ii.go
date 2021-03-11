package main

import (
	"fmt"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre_tail := dummy
	var next_head *ListNode = nil
	cnt := 0
	for ; head != nil; head = head.Next {
		cnt++
		if cnt == left {
			// start reversing
			var pre *ListNode = nil
			cur := head
			var reversed_head *ListNode = nil
			reversed_tail := head
			for cur != nil {
				tmp := cur.Next
				cur.Next = pre
				pre = cur
				cur = tmp

				if cnt == right {
					reversed_head = pre
					next_head = tmp
					break
				}
				cnt++
			}

			pre_tail.Next = reversed_head
			reversed_tail.Next = next_head

			break
		}
		pre_tail = head
	}

	return dummy.Next
}

func main() {
	head := Build([]int{1})
	fmt.Printf("reversed: %s\n", PrettyFormat(reverseBetween(head, 1, 1)))
}
