// Copyright (c) 2018 cloud-spin
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

package queueimpl7

import (
	"fmt"
	"sync"
	"testing"
)

type SafeQueue struct {
	q Queueimpl7
	m sync.Mutex
}

func (s *SafeQueue) Len() int {
	s.m.Lock()
	defer s.m.Unlock()

	return s.q.Len()
}

func (s *SafeQueue) Push(v interface{}) {
	s.m.Lock()
	defer s.m.Unlock()

	s.q.Push(v)
}

func (s *SafeQueue) Pop() (interface{}, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	return s.q.Pop()
}

func (s *SafeQueue) Front() (interface{}, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	return s.q.Front()
}

func TestSafeQueue(t *testing.T) {
	var q SafeQueue

	q.Push(1)
	q.Push(2)

	for v, ok := q.Pop(); ok; v, ok = q.Pop() {
		fmt.Println(v)
	}
}

func TestQueueImpl7NewQueueShouldReturnInitiazedInstanceOfQueue(t *testing.T) {
	ts := make([]int, 0, 1)

	ts = append(ts, 1)
	fmt.Println(cap(ts)) // Slice has 1 item; output: 1

	ts = append(ts, 1)
	fmt.Println(cap(ts)) // Slice has 2 items; output: 2

	ts = append(ts, 1)
	fmt.Println(cap(ts)) // Slice has 3 items; output: 4

	ts = append(ts, 1)
	ts = append(ts, 1)
	fmt.Println(cap(ts)) // Slice has 5 items; output: 8

	ts = append(ts, 1)
	ts = append(ts, 1)
	ts = append(ts, 1)
	ts = append(ts, 1)
	fmt.Println(cap(ts)) // Slice has 9 items; output: 16

	// q := New()

	// // Push 16 items to fill the first dynamic slice (sized 16).
	// for i := 1; i <= 16; i++ {
	// 	q.Push(i)
	// }

	// // Push an additional 128 items to fill the first full sized slice (sized 128).
	// for i := 1; i <= 128; i++ {
	// 	q.Push(i)
	// }

	// // Push 1 extra item that causes the creation of a new 128 sized slice to store this value,
	// // adding a total of 145 items to the queue.
	// q.Push(17)

	// // Pops the first 143 items to release the first dynamic slice (sized 16) and
	// // 127 items from the first full sized slice (sized 128).
	// for i := 1; i <= 143; i++ {
	// 	q.Pop()
	// }

	// // As unsafe.Sizeof (https://golang.org/pkg/unsafe/#Sizeof) doesn't consider the length of slices,
	// // we need to manually calculate the memory used by the internal slices.
	// var internalsliceType interface{}
	// fmt.Println(fmt.Sprintf("%d bytes", unsafe.Sizeof(q)+(unsafe.Sizeof(internalsliceType) /* bytes per slice position */ *(127 /* head slice unused positions */ +127 /* tail slice unused positions */))))

	// // Output for a 64bit system (Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz): 4072 bytes

	// if q == nil {
	// 	t.Error("Expected: new instance of queue; Got: nil")
	// }
}

func TestQueueImpl7WithZeroValuShouldReturnReadyToUseQueue(t *testing.T) {
	var q Queueimpl7
	q.Push(1)
	q.Push(2)

	v, ok := q.Pop()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = q.Pop()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	_, ok = q.Pop()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

func TestQueueImpl7WithNilValuesShouldReturnAllValuesInOrder(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(nil)
	q.Push(2)
	q.Push(nil)

	v, ok := q.Pop()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = q.Pop()
	if !ok || v != nil {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = q.Pop()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = q.Pop()
	if !ok || v != nil {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	_, ok = q.Pop()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

func TestQueueImpl7PutGetFrontShouldRetrieveAllElementsInOrder(t *testing.T) {
	tests := map[string]struct {
		putCount       []int
		getCount       []int
		remainingCount int
	}{
		"Test 1 item": {
			putCount:       []int{1},
			getCount:       []int{1},
			remainingCount: 0,
		},
		"Test 100 items": {
			putCount:       []int{100},
			getCount:       []int{100},
			remainingCount: 0,
		},
		"Test 1000 items": {
			putCount:       []int{1000},
			getCount:       []int{1000},
			remainingCount: 0,
		},
		"Test sequence 1": {
			putCount:       []int{1, 2, 100, 101},
			getCount:       []int{1, 2, 100, 101},
			remainingCount: 0,
		},
		"Test sequence 2": {
			putCount:       []int{10, 1},
			getCount:       []int{1, 10},
			remainingCount: 0,
		},
		"Test sequence 3": {
			putCount:       []int{101, 101},
			getCount:       []int{100, 101},
			remainingCount: 1,
		},
		"Test sequence 4": {
			putCount:       []int{1000, 1000, 1001},
			getCount:       []int{10, 10, 1},
			remainingCount: 2980,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			q := New()
			lastPut := 0
			lastGet := 0
			var ok bool
			var v interface{}
			for count := 0; count < len(test.getCount); count++ {
				for i := 1; i <= test.putCount[count]; i++ {
					lastPut++
					q.Push(lastPut)
					if v, ok = q.Front(); !ok || v != lastGet+1 {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
				}

				for i := 1; i <= test.getCount[count]; i++ {
					lastGet++
					v, ok = q.Front()
					if !ok || v.(int) != lastGet {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
					v, ok = q.Pop()
					if !ok || v.(int) != lastGet {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
				}
			}

			if q.Len() != test.remainingCount {
				t.Errorf("Expected: %d; Got: %d", test.remainingCount, q.Len())
			}

			if test.remainingCount > 0 {
				if v, ok = q.Front(); !ok || v == nil {
					t.Error("Expected: non-empty queue; Got: empty")
				}
			} else {
				if v, ok = q.Front(); ok || v != nil {
					t.Error("Expected: empty queue; Got: non-empty")
				}
			}

			for i := 1; i <= test.remainingCount; i++ {
				lastGet++

				if v, ok = q.Front(); !ok || v.(int) != lastGet {
					t.Errorf("Expected: %d; Got: %d", lastGet, v)
				}
				v, ok = q.Pop()
				if !ok || v.(int) != lastGet {
					t.Errorf("Expected: %d; Got: %d", lastGet, v)
				}
			}
			if v, ok = q.Front(); ok || v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
			if v, ok = q.Pop(); ok || v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
			if v, ok = q.Front(); ok || v != nil {
				t.Error("Expected: empty queue; Got: non-empty")
			}
			if q.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, q.Len())
			}
		})
	}
}
