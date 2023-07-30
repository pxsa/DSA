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

func TestPrint(t *testing.T) {
	l := &LinkedList{}

	l.Append("one")
	l.Append("two")

	res := l.Print()
	expected_answer := "<Node #1>: one\n<Node #2>: two\n"

	if res != expected_answer {
		t.Errorf("expected %s but recieved %s", expected_answer, res)
	}
	
}

func TestPrepend(t *testing.T) {
	l := &LinkedList{}

	l.Prepend("one")

	if l.Len() != 1 {
		t.Errorf("expected %v but got %v instead.", 1, l.Len())
	}

	l.Prepend("two")
	l.Prepend("three")

	if l.Len() != 3 {
		t.Errorf("expected %v but got %v instead.", 3, l.Len())
	}

	res, err := l.GetNode(1)
	if err != nil {
		t.Error(err)
	} else if res.Data != "two" {
		t.Errorf("expected %v but got %v instead.", "two", res.Data)
	}

	res, err = l.GetNode(2)
	if err != nil {
		t.Error(err)
	} else if res.Data != "one" {
		t.Errorf("expected %v but got %v instead.", "one", res.Data)
	}

	res, err = l.GetNode(2)
	if err != nil {
		t.Error(err)
	} else if res.Data != "one" {
		t.Errorf("expected %v but got %v instead.", "one", res.Data)
	}
}

func TestRemove(t *testing.T) {
	l := &LinkedList{}

	err := l.Remove(2)
	if err == nil {
		t.Error("expected an error but didn't recieve any (list is empty)")
	}

	l.Append("one")
	l.Prepend("two")
	l.Append("three")

	err = l.Remove(1)
	if err != nil {
		t.Error(err)
	}

	if l.Len() != 2 {
		t.Errorf("expected %v but got %v instead.", 2, l.Len())
	}

	res, err := l.GetNode(1)
	if err != nil {
		t.Error(err)
	} else if res.Data != "three" {
		t.Errorf("expected %v but got %v instead.", "three", res.Data)
	}

	err = l.Remove(12)
	if err == nil {
		t.Error("expected an error but didn't recieve any (index out of range)")
	}
}

func TestInse(t *testing.T) {
	l := &LinkedList{}

	l.Insert("zero", 2)
	l.Insert("one", 0)
	l.Insert("two", 0)
	l.Insert("three", 1)
	l.Insert("four", 1)

	if l.Len() != 4 {
		t.Errorf("expected %v but got %v instead.", 4, l.Len())
	}

	res, err := l.GetNode(1)
	if err != nil {
		t.Error(err)
	} else if res.Data != "four" {
		t.Errorf("expected %v but got %v instead.", "four", res.Data)
	}
}