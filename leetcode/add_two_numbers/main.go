package main

import "fmt"

func main() {
	node := addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}})
	fmt.Println(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}
		sum := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
