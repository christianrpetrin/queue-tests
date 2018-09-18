// Copyright (c) 2018 Christian R. Petrin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package queueimpl3 implements an unbounded, dynamically growing FIFO queue.
// Internally, queue store the values in fixed sized arrays that are linked using a singly linked list.
// This implementation tests the queue performance when controlling the length and current positions in the arrays using simple local variables instead of relying on the builtin len and append functions.
// This implementation also doesn't accept nil values on Push making it a bit easier to use the struct as only one value is returned by Front and Pop functions.
package queueimpl3

import "errors"

const (
	// internalArraySize holds the size of each internal array.
	internalArraySize = 64
)

// QueueImpl2 represents an unbounded, dynamically growing FIFO queue.
// The zero value for queue is an empty queue ready to use.
type QueueImpl2 struct {
	// Head points to the first node of the linked list.
	head *Node

	// Tail points to the last node of the linked list.
	// In an empty queue, head and tail points to the same node.
	tail *Node

	// Hp is the index pointing to the current first element in the queue (i.e. first element added in the current queue values).
	hp int

	// Tp is the index pointing to the current last element in the queue (i.e. last element added in the current queue values).
	tp int

	// Len holds the current queue values length.
	len int
}

// Node represents a queue node.
// Each node holds an array of user managed values.
type Node struct {
	// v holds the list of user added values in this node.
	v []interface{}

	// n points to the next node in the linked list.
	n *Node
}

// New returns an initialized queue.
func New() *QueueImpl2 {
	return new(QueueImpl2).Init()
}

// Init initializes or clears queue q.
func (q *QueueImpl2) Init() *QueueImpl2 {
	q.head = newNode()
	q.tail = q.head
	q.hp = 0
	q.tp = 0
	q.len = 0
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *QueueImpl2) Len() int { return q.len }

// Front returns the first element of list l or nil if the list is empty.
// If the queue is empty, nil is returned.
func (q *QueueImpl2) Front() interface{} {
	if q.len == 0 {
		return nil
	}

	return q.head.v[q.hp]
}

// Push adds a value to the queue.
func (q *QueueImpl2) Push(v interface{}) error {
	if v == nil {
		return errors.New("Cannot add nil value")
	}

	q.tail.v[q.tp] = v
	q.tp++
	q.len++

	if q.tp >= internalArraySize {
		q.tail.n = newNode()
		q.tail = q.tail.n
		q.tp = 0
	}

	return nil
}

// Pop retrieves and removes the next element from the queue.
// If the queue is empty, nil is returned.
func (q *QueueImpl2) Pop() interface{} {
	if q.len == 0 {
		return nil
	}

	v := q.head.v[q.hp]
	q.head.v[q.hp] = nil // Avoid memory leaks
	q.hp++
	q.len--
	if q.hp >= internalArraySize {
		q.head = q.head.n
		q.hp = 0
	}

	return v
}

// newNode returns an initialized node.
func newNode() *Node {
	return &Node{
		v: make([]interface{}, internalArraySize),
		n: nil,
	}
}
