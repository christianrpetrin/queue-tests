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

// Package queueimpl1 implements an unbounded, dynamically growing FIFO queue.
// This implementation tests the queue performance when storing the values in a simple slice.
// Pop removes the first slice element.
package queueimpl1

// Queueimpl1 represents an unbounded, dynamically growing FIFO queue.
type Queueimpl1 struct {
	// The queue values.
	v []interface{}
}

// New returns an initialized queue.
func New() *Queueimpl1 {
	return new(Queueimpl1).Init()
}

// Init initializes or clears queue q.
func (q *Queueimpl1) Init() *Queueimpl1 {
	q.v = make([]interface{}, 0)
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *Queueimpl1) Len() int { return len(q.v) }

// Front returns the first element of list l or nil if the list is empty.
// The second, bool result indicates whether a valid value was returned;
//   if the queue is empty, false will be returned.
// The complexity is O(1).
func (q *Queueimpl1) Front() (interface{}, bool) {
	if len(q.v) == 0 {
		return nil, false
	}

	return q.v[0], true
}

// Push adds a value to the queue.
// The complexity is amortized O(1).
func (q *Queueimpl1) Push(v interface{}) {
	q.v = append(q.v, v)
}

// Pop retrieves and removes the next element from the queue.
// The second, bool result indicates whether a valid value was returned; if the queue is empty, false will be returned.
// If the queue is empty, nil is returned.
// The complexity is amortized O(1).
func (q *Queueimpl1) Pop() (interface{}, bool) {
	if len(q.v) == 0 {
		return nil, false
	}

	v := q.v[0]
	q.v[0] = nil // Avoid memory leaks
	q.v = q.v[1:]

	return v, true
}
