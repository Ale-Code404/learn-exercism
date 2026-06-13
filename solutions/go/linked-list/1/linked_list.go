package linkedlist

import (
	"errors"
)

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

func NewList(elements ...any) *List {
	var list List

	for _, element := range elements {
		list.Push(element)
	}

	return &list
}

func (n *Node) Next() *Node {
	return n.Right
}

func (n *Node) Prev() *Node {
	return n.Left
}

func (l *List) Unshift(v any) {
	if l.Head == nil {
		node := &Node{v, nil, nil}
		l.Head = node
		l.Tail = node
		return
	}

	node := &Node{v, nil, l.Head}
	l.Head.Left = node

	l.Head = node
}

func (l *List) Push(v any) {
	if l.Tail == nil {
		node := &Node{v, nil, nil}
		l.Tail = node
		l.Head = node
		return
	}

	node := &Node{v, l.Tail, nil}
	l.Tail.Right = node

	l.Tail = node
}

func (l *List) Shift() (any, error) {
	if l.Head == nil {
		return nil, errors.New("There's no elements in the list")
	}

	first := l.Head

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Head = l.Head.Right
		l.Head.Left = nil
	}

	return first.Value, nil
}

func (l *List) Pop() (any, error) {
	if l.Tail == nil {
		return nil, errors.New("There's no elements in the list")
	}

	last := l.Tail

	if l.Tail == l.Head {
		l.Tail = nil
		l.Head = nil
	} else {
		l.Tail = l.Tail.Left
		l.Tail.Right = nil
	}

	return last.Value, nil
}

func (l *List) Reverse() {
	if l.Tail == nil {
		return
	}

	if l.Tail == l.Head {
		return
	}

	times := 0
	loops := int(l.Count() / 2)

	next := l.Head
	prev := l.Tail

	for times < loops {
		aux := next.Value
		next.Value = prev.Value
		prev.Value = aux

		next = next.Next()
		prev = prev.Prev()
		times++
	}
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Count() int {
	if l.Head == nil {
		return 0
	}

	length := 0
	for cursor := l.Head; cursor != nil; cursor = cursor.Next() {
		length++
	}

	return length
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	if ll.Head == nil {
		return false
	}

	var target *Node
	for cursor := ll.Head; cursor != nil; cursor = cursor.Next() {
		if cursor.Value != v {
			continue
		}

		target = cursor
		break
	}

	if target == nil {
		return false
	}

	if ll.Head == ll.Tail {
		ll.Tail = nil
		ll.Head = nil
		return true
	}

	if target == ll.Head {
		_, err := ll.Shift()
		return err != nil
	}

	if target == ll.Tail {
		_, err := ll.Pop()
		return err != nil
	}

	target.Left.Right = target.Right
	target.Right.Left = target.Left

	return true
}
