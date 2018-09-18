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

// Package queueimpl5 implements an unbounded, dynamically growing FIFO queue.
// Internally, queue store the values in fixed sized arrays that are linked using a singly linked list.
// This implementation tests the queue performance when storing the values in a simple slice. Pop moves the current position to next one instead of removing the first element.
package queueimpl5

import "errors"

// QueueImpl5 represents an unbounded, dynamically growing FIFO queue.
// The zero value for queue is an empty queue ready to use.
type QueueImpl5 struct {
	// The queue values.
	v []interface{}

	// The current value index.
	pos int
}

// New returns an initialized queue.
func New() *QueueImpl5 {
	return new(QueueImpl5).Init()
}

// Init initializes or clears queue q.
func (q *QueueImpl5) Init() *QueueImpl5 {
	q.v = make([]interface{}, 0)
	q.pos = 0
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *QueueImpl5) Len() int { return len(q.v) - q.pos }

// Front returns the first element of list l or nil if the list is empty.
// If the queue is empty, nil is returned.
func (q *QueueImpl5) Front() interface{} {
	if q.pos >= len(q.v) || q.v[q.pos] == nil {
		return nil
	}

	return q.v[q.pos]
}

// Push adds a value to the queue.
func (q *QueueImpl5) Push(v interface{}) error {
	if v == nil {
		return errors.New("Cannot add nil value")
	}

	q.v = append(q.v, v)
	return nil
}

// Pop retrieves and removes the next element from the queue.
// If the queue is empty, nil is returned.
func (q *QueueImpl5) Pop() interface{} {
	if q.pos >= len(q.v) || q.v[q.pos] == nil {
		return nil
	}

	v := q.v[q.pos]
	q.v[q.pos] = nil // Avoid memory leaks
	q.pos++

	return v
}

// newNode returns an initialized array of interface{}.
func newNode() []interface{} {
	return make([]interface{}, 1)
}
