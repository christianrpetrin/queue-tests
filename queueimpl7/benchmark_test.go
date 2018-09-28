// Copyright (c) 2018 Christian R. Petrin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software out restriction, including out limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", OUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION  THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package queueimpl7

import (
	"strconv"
	"testing"
)

const (
	// count holds the number of items to add to the queue.
	count = 1000
)

var (
	// Used to store temp values, avoiding any compiler optimizations.
	tmp  interface{}
	tmp2 bool
)

func BenchmarkMaxFirstSliceSize(b *testing.B) {
	tests := []struct {
		name string
		size int
	}{
		{size: 1},
		{size: 2},
		{size: 4},
		{size: 8},
		{size: 16},
		{size: 32},
		{size: 64},
		{size: 128},
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxFirstSliceSize = test.size
				q := New()

				for i := 0; i < count; i++ {
					q.Push(n)
				}
				for tmp, tmp2 = q.Pop(); tmp2; tmp, tmp2 = q.Pop() {
				}
			}
		})
	}
}

func BenchmarkMaxSubsequentSliceSize(b *testing.B) {
	tests := []struct {
		name string
		size int
	}{
		{size: 1},
		{size: 10},
		{size: 20},
		{size: 30},
		{size: 40},
		{size: 50},
		{size: 60},
		{size: 70},
		{size: 80},
		{size: 90},
		{size: 100},
		{size: 200},
		{size: 2},
		{size: 4},
		{size: 8},
		{size: 16},
		{size: 32},
		{size: 64},
		{size: 128},
		{size: 256},
		{size: 512},
		{size: 1024},
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxInternalSliceSize = test.size
				q := New()

				for i := 0; i < count; i++ {
					q.Push(n)
				}
				for tmp, tmp2 = q.Pop(); tmp2; tmp, tmp2 = q.Pop() {
				}
			}
		})
	}
}
