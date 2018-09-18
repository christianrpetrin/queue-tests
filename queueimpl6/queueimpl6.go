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

// Package QueueImpl6 implements a bounded FIFO queue.
// This implementation tests the queue performance when storing the values in a simple slice. Pop removes the first slice element.
package queueimpl6

import "errors"

// QueueImpl6 represents an unbounded, dynamically growing FIFO queue.
// The zero value for queue is an empty queue ready to use.
type QueueImpl6 struct {
	// The queue values.
	v []interface{}
}

// New returns an initialized queue.
func New() *QueueImpl6 {
	return new(QueueImpl6).Init()
}

// Init initializes or clears queue q.
func (q *QueueImpl6) Init() *QueueImpl6 {
	q.v = make([]interface{}, 0)
	return q
}

// Len returns the number of elements of queue q.
func (q *QueueImpl6) Len() int { return len(q.v) }

// Front returns the first element of list l or nil if the list is empty.
// If the queue is empty, nil is returned.
func (q *QueueImpl6) Front() interface{} {
	if len(q.v) == 0 {
		return nil
	}

	return q.v[0]
}

// Push adds a value to the queue.
func (q *QueueImpl6) Push(v interface{}) error {
	if v == nil {
		return errors.New("Cannot add nil value")
	}

	q.v = append(q.v, v)
	return nil
}

// Pop retrieves and removes the next element from the queue.
// If the queue is empty, nil is returned.
func (q *QueueImpl6) Pop() interface{} {
	if len(q.v) == 0 {
		return nil
	}

	v := q.v[0]
	q.v[0] = nil // Avoid memory leaks
	q.v = q.v[1:]

	return v
}

// newNode returns an initialized array of interface{}.
func newNode() []interface{} {
	return make([]interface{}, 1)
}
