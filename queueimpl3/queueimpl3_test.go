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

package queueimpl3

import (
	"testing"
)

func TestQueueImpl3NewQueueShouldReturnInitiazedInstanceOfQueue(t *testing.T) {
	q := New()

	if q == nil {
		t.Error("Expected: new instance of queue; Got: nil")
	}
}

func TestQueueImpl3WithNilValuesShouldReturnError(t *testing.T) {
	q := New()

	if err := q.Push(nil); err == nil {
		t.Error("Expected: error; Got: success")
	}
}

func TestQueueImpl3PutGetFrontShouldRetrieveAllElementsInOrder(t *testing.T) {
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
			putCount:       []int{2},
			getCount:       []int{2},
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
			var v interface{}
			for count := 0; count < len(test.getCount); count++ {
				for i := 1; i <= test.putCount[count]; i++ {
					lastPut++
					q.Push(lastPut)
					if v = q.Front(); v != lastGet+1 {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
				}

				for i := 1; i <= test.getCount[count]; i++ {
					lastGet++
					if v = q.Front(); v.(int) != lastGet {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
					if v = q.Pop(); v.(int) != lastGet {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
				}
			}

			if q.Len() != test.remainingCount {
				t.Errorf("Expected: %d; Got: %d", test.remainingCount, q.Len())
			}

			if test.remainingCount > 0 {
				if v = q.Front(); v == nil {
					t.Error("Expected: non-empty queue; Got: empty")
				}
			} else {
				if v = q.Front(); v != nil {
					t.Error("Expected: empty queue; Got: non-empty")
				}
			}

			for i := 1; i <= test.remainingCount; i++ {
				lastGet++

				if v = q.Front(); v.(int) != lastGet {
					t.Errorf("Expected: %d; Got: %d", lastGet, v)
				}
				if v = q.Pop(); v.(int) != lastGet {
					t.Errorf("Expected: %d; Got: %d", lastGet, v)
				}
			}
			if v = q.Front(); v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
			if v = q.Pop(); v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
			if v = q.Front(); v != nil {
				t.Error("Expected: empty queue; Got: non-empty")
			}
			if q.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, q.Len())
			}
		})
	}
}
