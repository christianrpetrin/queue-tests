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

// Package queueimpl4 implements an unbounded, dynamically growing FIFO queue.
// This implementation tests the queue performance when storing the "next" pointer as part of the values array instead of using a separate field.
// The next element is stored in the last position of the internal array, which is a reserved position.
// A node is each internal array that is used to store the elements.
package queueimpl4

import "errors"

const (
	// internalArraySize holds the size of each internal array.
	internalArraySize = 64
	// internalArrayLinkPosition holds the last position of the internal array.
	internalArrayLinkPosition = 63
)

// QueueImpl4 represents an unbounded, dynamically growing FIFO queue.
// The zero value for queue is an empty queue ready to use.
type QueueImpl4 struct {
	// Head points to the first node of the linked list.
	head []interface{}

	// Tail points to the last node of the linked list.
	// In an empty queue, head and tail points to the same node.
	tail []interface{}

	// Hp is the index pointing to the current first element in the queue (i.e. first element added in the current queue values).
	hp int

	// Tp is the index pointing to the current last element in the queue (i.e. last element added in the current queue values).
	tp int

	// Len holds the current queue values length.
	len int
}

// New returns an initialized queue.
func New() *QueueImpl4 {
	return new(QueueImpl4).Init()
}

// Init initializes or clears queue q.
func (q *QueueImpl4) Init() *QueueImpl4 {
	q.head = newNode()
	q.tail = q.head
	q.hp = 0
	q.tp = 0
	q.len = 0
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *QueueImpl4) Len() int { return q.len }

// Front returns the first element of list l or nil if the list is empty.
// If the queue is empty, nil is returned.
// The complexity is O(1).
func (q *QueueImpl4) Front() interface{} {
	if q.head[q.hp] == nil {
		return nil
	}

	return q.head[q.hp]
}

// Push adds a value to the queue.
// The complexity is O(1).
func (q *QueueImpl4) Push(v interface{}) error {
	if v == nil {
		return errors.New("Cannot add nil value")
	}

	q.tail[q.tp] = v
	q.tp++
	q.len++
	if q.tp >= internalArrayLinkPosition {
		node := newNode()
		q.tail[q.tp] = node
		q.tail = node
		q.tp = 0
	}

	return nil
}

// Pop retrieves and removes the next element from the queue.
// If the queue is empty, nil is returned.
// The complexity is O(1).
func (q *QueueImpl4) Pop() interface{} {
	if q.len == 0 {
		return nil
	}

	v := q.head[q.hp]
	q.head[q.hp] = nil // Avoid memory leaks
	q.hp++
	q.len--
	if q.hp >= internalArrayLinkPosition {
		q.head = q.head[q.hp].([]interface{})
		q.hp = 0
	}

	return v
}

// newNode returns an initialized array of interface{}.
func newNode() []interface{} {
	return make([]interface{}, internalArraySize)
}
