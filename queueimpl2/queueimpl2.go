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

// Package queueimpl2 implements an unbounded, dynamically growing FIFO queue.
// This implementation tests the queue performance when storing the values in a simple
// slice. Pop moves the current position to next one instead of removing the first element.
package queueimpl2

// Queueimpl2 represents an unbounded, dynamically growing FIFO queue.
type Queueimpl2 struct {
	// The queue values.
	v []interface{}

	// The current value index.
	pos int
}

// New returns an initialized queue.
func New() *Queueimpl2 {
	return new(Queueimpl2).Init()
}

// Init initializes or clears queue q.
func (q *Queueimpl2) Init() *Queueimpl2 {
	q.v = make([]interface{}, 0)
	q.pos = 0
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *Queueimpl2) Len() int { return len(q.v) - q.pos }

// Front returns the first element of list l or nil if the list is empty.
// The second, bool result indicates whether a valid value was returned;
// if the queue is empty, false will be returned.
// The complexity is O(1).
func (q *Queueimpl2) Front() (interface{}, bool) {
	if q.pos >= len(q.v) {
		return nil, false
	}

	return q.v[q.pos], true
}

// Push adds a value to the queue.
// The complexity is amortized O(1).
func (q *Queueimpl2) Push(v interface{}) {
	q.v = append(q.v, v)
}

// Pop retrieves and removes the next element from the queue.
// The second, bool result indicates whether a valid value was returned;
//   if the queue is empty, false will be returned.
// The complexity is O(1).
func (q *Queueimpl2) Pop() (interface{}, bool) {
	if q.pos >= len(q.v) {
		return nil, false
	}

	v := q.v[q.pos]
	q.v[q.pos] = nil // Avoid memory leaks
	q.pos++

	return v, true
}
