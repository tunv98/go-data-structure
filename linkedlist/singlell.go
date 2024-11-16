package main

import "fmt"

type SingleLL[T any] struct {
	head   *Node[T]
	length int
}

func NewSingleLL[T any]() *SingleLL[T] {
	return &SingleLL[T]{}
}

// AddElementToHead adds a new element to the head of the linked list.
//
// val: the value of the element to be added.
func (l *SingleLL[T]) AddElementToHead(val T) {
	newNode := NewNode(val)
	newNode.Next = l.head
	l.head = newNode
	l.length++
}

// AddElementToTail adds a new element to the tail of the single linked list.
//
// val T
func (l *SingleLL[T]) AddElementToTail(val T) {
	newNode := NewNode(val)
	l.length++
	if l.head == nil {
		l.head = newNode
		return
	}
	tail := l.head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = newNode
}

// RemoveElementFromHead removes the element at the head of the SingleLL.
//
// No parameters.
// No return types.
func (l *SingleLL[T]) RemoveElementFromHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.Next
	l.length--
}

// RemoveElementFromTail removes the last element from the linked list.
//
// No parameters.
func (l *SingleLL[T]) RemoveElementFromTail() {
	if l.head == nil {
		return
	}
	if l.head.Next == nil {
		l.head = nil
		l.length--
		return
	}
	tail := l.head
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	tail.Next = nil
	l.length--
}

// Length returns the length of the SingleLL.
// No parameters.
// Returns an integer.
func (l *SingleLL[T]) Length() int {
	return l.length
}

// Reserve reverses the linked list.
//
// No parameters. No return type.
func (l *SingleLL[T]) Reserve() {
	var prev, next *Node[T]
	curr := l.head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *SingleLL[T]) Print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}
