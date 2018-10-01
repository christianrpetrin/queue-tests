// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package event provides an infinitely buffered double-ended queue of events.
package bcmills

// Deque is an infinitely buffered double-ended queue of events. The zero value
// is usable, but a Deque value must not be copied.
type Deque struct {
	// mu    sync.Mutex
	// cond  sync.Cond     // cond.L is lazily initialized to &Deque.mu.
	front []interface{} // LIFO.
	back  []interface{} // FIFO; used only when front is not empty.
}

// func (q *Deque) lockAndInit() {
// 	q.mu.Lock()
// 	if q.cond.L == nil {
// 		q.cond.L = &q.mu
// 	}
// }

// NextEvent implements the screen.EventDeque interface.
func (q *Deque) NextEvent() interface{} {
	// q.lockAndInit()
	// defer q.mu.Unlock()

	for len(q.front) == 0 {
		return nil
		//q.cond.Wait()
	}
	i := len(q.front) - 1
	event := q.front[i]
	q.front[i] = nil
	q.front = q.front[:i]

	if len(q.front) == 0 {
		for n := len(q.back); n > 0; n-- {
			q.front = append(q.front, q.back[n-1])
			q.back[n-1] = nil
		}
		q.back = q.back[:0]
	}
	return event
}

// Send implements the screen.EventDeque interface.
func (q *Deque) Send(event interface{}) {
	// q.lockAndInit()
	// defer q.mu.Unlock()

	if len(q.front) == 0 && len(q.back) == 0 {
		q.front = append(q.front, event)
	} else {
		q.back = append(q.back, event)
	}
	// q.cond.Signal()
}

// SendFirst implements the screen.EventDeque interface.
func (q *Deque) SendFirst(event interface{}) {
	// q.lockAndInit()
	// defer q.mu.Unlock()

	q.front = append(q.front, event)
	// q.cond.Signal()
}
