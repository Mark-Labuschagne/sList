// BSD 2-Clause License

// Copyright (c) 2021, Mark-Labuschagne
// All rights reserved.

package linkedList

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrEmptyList       = errors.New("Empty list")
	ErrNodeNotFound    = errors.New("Node does not exist")
	ErrMismatchedTypes = errors.New("Mismatched types")
)

type Node struct {
	Next *Node
	Data interface{}
}

type List struct {
	Head  *Node
	Tail  *Node
	Typed bool
}

/*
Instantiate a singly linked list. To make the list typed, set argument typed to true, else set typed to false.
If you set typed to false, you can ignore error checking on (List).Insert().
*/
func CreateList(typed bool) *List {
	return &List{
		Head:  nil,
		Typed: typed,
	}
}

// Insert a node into the list.
func (l *List) Insert(d interface{}) error {
	newNode := &Node{
		Data: d,
	}

	if l.Head == nil {
		l.Head = newNode
	} else {
		if l.Typed && fmt.Sprintf("%T", l.Head.Data) != fmt.Sprintf("%T", d) {
			return errors.Wrap(ErrMismatchedTypes, "Cannot insert value of type ("+fmt.Sprintf("%T", d)+
				") into list having type ("+fmt.Sprintf("%T", l.Head.Data)+")")
		}

		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}

	return nil
}

// Display the entries in the list.
func (l *List) Display() error {
	err := l.checkEmpty()
	if err != nil {
		return err
	}

	var (
		first = l.Head
		i     = 1
	)
	fmt.Printf("Node %3v | Value %5v | Next Address %v\n", "", "", "")
	for first.Next != nil {
		fmt.Printf("%-10v %-13v %p\n", i, first.Data, first.Next)
		first = first.Next
		i++
	}
	fmt.Printf("%-10v %-13v <nil>\n", i, first.Data)

	return nil
}

// Remove a node from the list where (Node).Data == val.
func (l *List) RemoveNode(val interface{}) error {
	err := l.checkEmpty()
	if err != nil {
		return err
	}

	var (
		prev  *Node
		broke bool
		curr  = l.Head
	)
	for curr.Next != nil {
		if curr.Data == val {
			l.reorder(prev, curr.Next)
			broke = true
			break
		}
		prev = curr
		curr = curr.Next
	}
	if !broke && curr.Data == val {
		l.reorder(prev, curr.Next)
	} else {
		return ErrNodeNotFound
	}

	return nil
}

func (l *List) reorder(prev, next *Node) {
	curr := l.Head
	if prev == nil {
		l.Head = curr.Next
		return
	}
	for curr.Next != nil {
		if &curr.Next == &prev.Next {
			curr.Next = next
			break
		}
		curr = curr.Next
	}
}

// Remove duplicate values from the list.
// Returns an error if the list is empty.
func (l *List) RemoveDuplicates() error {
	err := l.checkEmpty()
	if err != nil {
		return err
	}

	var (
		prev *Node
		curr = l.Head
		dMap = make(map[interface{}]bool)
	)
	for curr.Next != nil {
		if dMap[curr.Data] {
			curr = curr.Next
			prev.Next = curr
		} else {
			dMap[curr.Data] = true
			prev = curr
			curr = curr.Next
		}
	}
	if dMap[curr.Data] {
		prev.Next = l.Tail
	}

	return nil
}

var (
	currN  *Node
	called bool
)

/*
Range over the list. Returns each value of each node in the list.

Usage:
	for end, val := l.Range(); !end; end, val = l.Range() {
		// Do whatever
	}
*/
func (l List) Range() (end bool, val interface{}) {
	if !called {
		if l.checkEmpty() != nil {
			return true, nil
		}

		currN = l.Head
		called = true
	} else {
		if currN.Next == nil {
			currN = &Node{}
			called = false

			return true, currN.Data
		}

		currN = currN.Next
	}

	return false, currN.Data
}

func (l List) checkEmpty() error {
	if l.Head == nil {
		return ErrEmptyList
	}

	return nil
}
