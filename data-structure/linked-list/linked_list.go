package main

import "fmt"

type Node struct {
	key  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(k int) {
	if l.head == nil {
		l.head = &Node{key: k}
	} else {
		l.head.Insert(k)
	}
}

func (n *Node) Insert(k int) {
	if n.next == nil {
		n.next = &Node{key: k}
		return
	} else {
		n.next.Insert(k)
	}
}

func (l *LinkedList) Print() {
	fmt.Println("\nDisplay the linked list :")
	if l.head == nil {
		fmt.Println("Linked list is empty")
	}
	node := l.head
	for node != nil {
		fmt.Printf("%v =>", node.key)
		node = node.next
	}
}

func (list *LinkedList) Reverse() {
	node := list.head

	var next *Node
	for node != nil {
		node, next, node.next = node.next, node, next
	}

	list.head.next = nil

	list.head = next
}

func main() {
	ll := &LinkedList{}
	ll.Insert(4)
	ll.Insert(14)
	ll.Insert(314)
	ll.Insert(1214)
	ll.Insert(5)

	ll.Print()
	ll.Reverse()
	ll.Print()
}
