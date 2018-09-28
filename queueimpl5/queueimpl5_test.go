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

package queueimpl5

import (
	"testing"
)

func TestQueueImpl1NewQueueShouldReturnInitiazedInstanceOfQueue(t *testing.T) {
	q := New()

	if q == nil {
		t.Error("Expected: new instance of queue; Got: nil")
	}
}

func TestQueueImpl1WithNilValuesShouldReturnAllValuesInOrder(t *testing.T) {
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

func TestQueueImpl1PutGetFrontShouldRetrieveAllElementsInOrder(t *testing.T) {
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
