package doubly

import (
	"testing"
)

var Tests = []struct {
	Name string
	Data any
	Length int
} {
	{"first", "one", 1},
	{"second", "two", 2},
}

func TestPrepend(t *testing.T) {
	l := &DoublyLL{}

	if l.Len() != 0 {
		t.Errorf("expected %v but got %v", 0, l.Len())
	}

	for _, test := range Tests {
		t.Run(test.Name, func(t *testing.T) {
			l.Prepend(test.Data)
			if l.Len() != test.Length {
				t.Errorf("expected %v but got %v", test.Length, l.Len())
			}
		})
	}
}

func TestAppend(t *testing.T) {
	l := &DoublyLL{}

	if l.Len() != 0 {
		t.Errorf("expected %v but got %v", 0, l.Len())
	}

	for _, test := range Tests {
		
		t.Run(test.Name, func(t *testing.T) {

			l.Append(test.Data)

			if l.Len() != test.Length {
				t.Errorf("expected %v but got %v", test.Length, l.Len())
			}

		})
	}
}