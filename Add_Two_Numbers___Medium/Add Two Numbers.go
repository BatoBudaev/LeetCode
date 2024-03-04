package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(value int) {
	newNode := &ListNode{value, nil}

	curr := l
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func printList(l *ListNode) {
	curr := l
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}

	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		carry = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 9 {
			sum %= 10
			carry += 1
		}

		curr.Val = sum

		if l1 != nil || l2 != nil || carry > 0 {
			node := new(ListNode)
			curr.Next = node
			curr = node
		}
	}

	return head
}

func main() {
	l1 := &ListNode{2, nil}
	l1.add(3)
	l1.add(4)
	printList(l1)

	l2 := &ListNode{5, nil}
	l2.add(6)
	l2.add(4)
	printList(l2)

	l3 := addTwoNumbers(l1, l2)
	printList(l3)
}
