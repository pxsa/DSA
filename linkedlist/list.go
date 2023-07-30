package linkedlist

import "fmt"

// Node is a container used in single linked list
type Node struct {
	Data any
	Next *Node
}

// LinkedList represents the single linked list
type LinkedList struct {
	Head *Node
}

// Append appends a new node to the linkedlist
func (l *LinkedList) Append(data any) {
	
	new := &Node{Data: data, Next: nil}
	
	if l.Head == nil {
		l.Head = new
		return
	}
	
	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = new
}

// Print prints all nodes of linkedlist
func (l *LinkedList) Print() {
	cur := l.Head
	var index int
	for cur != nil {
		index += 1
		fmt.Printf("<Node #%d>: %v", index, cur.Data)
		cur = cur.Next
	}
}

// Len returns the length of the linkedlist
func (l *LinkedList) Len() int {
	var count int
	cur := l.Head

	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}