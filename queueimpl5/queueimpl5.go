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
// Internally, queue store the values in fixed sized slices that are linked using
// a singly linked list.
// This implementation tests the queue performance when storing the "next" pointer
// as part of the values slice instead of having it as a separate field.
// The next element is stored in the last position of the internal slice, which is a reserved position.
// A node is each internal slice that is used to store the elements.
package queueimpl5

const (
	// internalSliceSize holds the size of each internal slice.
	internalSliceSize = 128

	// internalSliceLastPosition holds the last position of the internal slice.
	internalSliceLastPosition = 127
)

// Queueimpl5 represents an unbounded, dynamically growing FIFO queue.
// The zero value for queue is an empty queue ready to use.
type Queueimpl5 struct {
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
func New() *Queueimpl5 {
	return new(Queueimpl5).Init()
}

// Init initializes or clears queue q.
func (q *Queueimpl5) Init() *Queueimpl5 {
	n := newNode()
	q.head = n
	q.tail = n
	q.hp = 0
	q.tp = 0
	q.len = 0
	return q
}

// Len returns the number of elements of queue q.
// The complexity is O(1).
func (q *Queueimpl5) Len() int { return q.len }

// Front returns the first element of list l or nil if the list is empty.
// The second, bool result indicates whether a valid value was returned;
//   if the queue is empty, false will be returned.
// The complexity is O(1).
func (q *Queueimpl5) Front() (interface{}, bool) {
	if q.len == 0 {
		return nil, false
	}

	return q.head[q.hp], true
}

// Push adds a value to the queue.
// The complexity is O(1).
func (q *Queueimpl5) Push(v interface{}) {
	if q.tp >= internalSliceLastPosition {
		n := newNode()
		q.tail[q.tp] = n
		q.tail = n
		q.tp = 0
	}

	q.tail[q.tp] = v
	q.len++
	q.tp++
}

// Pop retrieves and removes the next element from the queue.
// The second, bool result indicates whether a valid value was returned;
//   if the queue is empty, false will be returned.
// The complexity is O(1).
func (q *Queueimpl5) Pop() (interface{}, bool) {
	if q.len == 0 {
		return nil, false
	}

	v := q.head[q.hp]
	q.head[q.hp] = nil // Avoid memory leaks
	q.hp++
	q.len--

	if q.hp >= internalSliceLastPosition {
		h := q.head[q.hp].([]interface{})
		q.head[q.hp] = nil // Avoid memory leaks
		q.head = h
		q.hp = 0
	}

	return v, true
}

// newNode returns an initialized slice of interface{}.
func newNode() []interface{} {
	return make([]interface{}, internalSliceSize)
}
