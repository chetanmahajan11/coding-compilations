package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func testLL() {
	n1 := &Node{Val: 10}
	n2 := &Node{Val: 20}
	n3 := &Node{Val: 30}

	head := n1
	n1.Next = n2
	n2.Next = n3

	print(head)
	head = insertAtBeginning(head, 50)
	print(head)
	head = insertAtEnd(head, 40)
	print(head)
	// head = delete(head, 30)
	// print(head)
	// head = reverseLL(head)
	// print(head)
	middleNum := middle(head)
	fmt.Println("middle is: ", middleNum.Val)
}

func print(head *Node) {
	current := head
	fmt.Println("------------------------")
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("\n------------------------")
}

func insertAtBeginning(head *Node, val int) *Node {
	return &Node{Val: val, Next: head}
}

func insertAtEnd(head *Node, val int) *Node {
	newNode := &Node{Val: val}
	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}

func delete(head *Node, val int) *Node {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	current := head

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}

	return head
}

func reverseLL(head *Node) *Node {
	var current, prev *Node
	current = head

	for current != nil {
		next := current.Next // stores next
		current.Next = prev  // flip arrow
		prev = current       // shift prev
		current = next       // shift current
	}

	return prev
}

func middle(head *Node) *Node {
	var slow, fast *Node

	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	/*// for first middle of LL:
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}*/
	return slow
}
