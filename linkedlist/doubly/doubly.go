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
	}

	return length
}

// insert: prepend
func (l *DoublyLL) Prepend(data any) {
	// var new_node *Node
	new_node := new(Node)
	new_node.Data = data
	// new_node := &Node{Data: data}
	
	// l is empty
	l.Head = new_node
	l.Tail = new_node

	// l is not empty
	new_node.Next = l.Head
	l.Head = new_node
}

// search
// delete
// update