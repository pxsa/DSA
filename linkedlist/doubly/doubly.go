package doubly

type Node struct {
	Data any
	Next *Node
	Prev *Node
}

type DoublyLL struct {
	Head *Node
	Tail *Node
}

// len
func (l *DoublyLL) Len() int {
	var length int
	cur := l.Head

	for cur != nil {
		length += 1
		cur = cur.Next
	}

	return length
}

// insert: prepend
func (l *DoublyLL) Prepend(data any) {
	var new_node *Node = new(Node)
	new_node.Data = data

	// l is empty
	if l.Len() == 0 {
		l.Head = new_node
		l.Tail = new_node

	// l is not empty
	} else {
		l.Head.Prev = new_node
		new_node.Next = l.Head
		l.Head = new_node
	}
}

// insert: append
func (l *DoublyLL) Append(data any) {
	
	newNode := &Node{Data: data}

	if l.Len() == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
}

// search
// delete
// update