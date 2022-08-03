package singly

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// New constructs and returns an empty singly linked-list.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// IsEmpty returns true if the linked-list doesn't contain any nodes.
func (s *LinkedList[T]) IsEmpty() bool {
	return s.Size == 0
}

// First returns the first element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) First() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}
	return s.Head.Data, true
}

// Last returns the last element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) Last() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}
	return s.Tail.Data, true
}

// AddFirst adds a new node to the beginning of the list.
func (s *LinkedList[T]) AddFirst(data T) {
	s.Head = &Node[T]{Data: data, Next: s.Head}

	if s.Size == 0 {
		s.Tail = s.Head
	}

	s.Size++
}

// AddLast adds a new node to the end of the list.
func (s *LinkedList[T]) AddLast(data T) {
	n := &Node[T]{Data: data}

	if s.IsEmpty() {
		s.Head = n
	} else {
		s.Tail.Next = n
	}

	s.Tail = n

	s.Size++
}

// Add adds a new node to the given index in the list. It returns ErrInvalidIndex if the given index is out of bound.
func (s *LinkedList[T]) Add(data T, index int) error {
	if index < 0 || index > s.Size {
		return ErrInvalidIndex
	}

	if index == 0 {
		s.AddFirst(data)
		return nil
	}
	if index == s.Size {
		s.AddLast(data)
		return nil
	}

	var count int

	current := s.Head
	for ; current != nil; current = current.Next {
		count++
		if count == index {
			break
		}
	}

	n := Node[T]{Data: data, Next: current.Next}
	current.Next = &n

	s.Size++

	return nil
}

// RemoveFirst removes and returns the first element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) RemoveFirst() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.Head.Data

	s.Head = s.Head.Next
	s.Size--

	if s.IsEmpty() {
		s.Tail = nil
	}

	return val, true
}

// RemoveLast removes and returns the last element of the list. It returns false if the list empty.
func (s *LinkedList[T]) RemoveLast() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.Tail.Data

	if s.Size == 1 {
		s.Tail = nil
		s.Head = nil
		s.Size--

		return val, true
	}

	current := s.Head
	for ; current.Next.Next != nil; current = current.Next {
	}

	current.Next = nil
	s.Tail = current
	s.Size--

	return val, true
}

// Remove removes and returns the node in the given index. It returns ErrInvalidIndex and ErrEmptyList
// if the list is empty or the index is out of bound.
func (s *LinkedList[T]) Remove(index int) (val T, err error) {
	if s.IsEmpty() {
		return val, ErrEmptyList
	}

	if index < 0 || index >= s.Size {
		return val, ErrInvalidIndex
	}

	if index == 0 {
		val, _ := s.RemoveFirst()
		return val, nil
	}

	if index == s.Size-1 {
		val, _ := s.RemoveLast()
		return val, nil
	}

	var count int

	current := s.Head
	for ; current != nil; current = current.Next {
		count++
		if count == index {
			break
		}
	}

	val = current.Next.Data

	current.Next = current.Next.Next
	s.Size--

	return val, nil
}

// String returns the string representation of the list.
func (s *LinkedList[T]) String() string {
	var b strings.Builder

	b.WriteString("[ ")

	current := s.Head
	for ; current != nil; current = current.Next {
		b.WriteString(fmt.Sprint(current.Data))
		b.WriteString(" ")
	}

	b.WriteString("]")

	return b.String()
}

// ToSlice returns the linked-list as a slice.
func (s *LinkedList[T]) ToSlice() []T {
	r := make([]T, s.Size)

	for i, cur := 0, s.Head; cur != nil && i < len(r); i, cur = i+1, cur.Next {
		r[i] = cur.Data
	}

	return r
}
