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
	Tail *Node
}

// Append appends a new node to the linkedlist
func (l *LinkedList) Append(data any) {
	
	new := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = new
		l.Tail = new
		return
	}
	
	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = new
	l.Tail = new
}

// Prepend prepends a new node to the linkedlist
func (l *LinkedList) Prepend(data any) {

	new_node := &Node{
		Data: data,
		Next: l.Head,
	}

	l.Head = new_node
}

// Insert inserts a new node at a specific location
func (l *LinkedList) Insert(data any, index int) {
	if index == 0 {
		l.Prepend(data)
	} else if (index == l.Len()-1) {
		l.Append(data)
	} else if index < l.Len() {
		
		var curIndex int
		preNode := l.Head

		for curIndex != index-1 {
			curIndex++
			preNode = preNode.Next
		}

		newNode := &Node{Data: data, Next: preNode.Next}
		preNode.Next = newNode

	}
}

// RemoveFirst removes the first node from list
func (l *LinkedList) RemoveFirst() error {
	if l.Len() == 0 {
		return fmt.Errorf("list is empty")
	}

	l.Head = l.Head.Next

	if l.Len() == 1 {
		l.Tail = nil
	}

	return nil
}

// Remove removes the specified node from list
func (l *LinkedList) Remove(index int) error {

	if l.Len() == 0 {
		return fmt.Errorf("list is empty")
	}

	if index >= l.Len() {
		return fmt.Errorf("index out of range")
	}  

	cur := l.Head
	curIndex := 0
	for curIndex != index-1 {
		curIndex++
		cur = cur.Next
	}

	if index == l.Len() - 1 {
		l.Tail = cur
		cur.Next = nil
	} else {
		cur.Next = cur.Next.Next
	}

	return nil
}

// Print prints all nodes of linkedlist
func (l *LinkedList) Print() string {
	cur := l.Head
	var index int
	var res string

	for cur != nil {
		index += 1
		res += fmt.Sprintf("<Node #%d>: %v\n", index, cur.Data)
		cur = cur.Next
	}

	return res
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

func (l *LinkedList) GetNode(index int) (*Node, error) {
	curIndex := 0
	cur := l.Head

	if index >= l.Len() {
		return nil, fmt.Errorf("out of range")
	}

	if l.Len() == 0 {
		return nil, fmt.Errorf("list is empty")
	}

	for curIndex < index {
		cur = cur.Next
		curIndex += 1
	}

	return cur, nil
}