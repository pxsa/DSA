package linkedlist

import (
	"testing"
)

func TestAppend(t *testing.T) {
	// initialize a linked list
	l := &LinkedList{}

	l.Append("first node")
	if l.Len() != 1 {
		t.Errorf("expected to recieve %d but but recieve %d", 1, l.Len())
	}

	l.Append("seconde node")
	if l.Len() != 2 {
		t.Errorf("expected to recieve %d but but recieve %d", 2, l.Len())
	}
}