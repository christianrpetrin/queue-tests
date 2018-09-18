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

package queueimpl1

import (
	"testing"
)

func BenchmarkQueue1(b *testing.B) {
	benchmarkQueue(b, 1)
}

func BenchmarkQueue10(b *testing.B) {
	benchmarkQueue(b, 10)
}

func BenchmarkQueue20(b *testing.B) {
	benchmarkQueue(b, 20)
}

func BenchmarkQueue30(b *testing.B) {
	benchmarkQueue(b, 30)
}

func BenchmarkQueue40(b *testing.B) {
	benchmarkQueue(b, 40)
}

func BenchmarkQueue50(b *testing.B) {
	benchmarkQueue(b, 50)
}

func BenchmarkQueue60(b *testing.B) {
	benchmarkQueue(b, 60)
}

func BenchmarkQueue70(b *testing.B) {
	benchmarkQueue(b, 70)
}

func BenchmarkQueue80(b *testing.B) {
	benchmarkQueue(b, 80)
}

func BenchmarkQueue90(b *testing.B) {
	benchmarkQueue(b, 90)
}

func BenchmarkQueue100(b *testing.B) {
	benchmarkQueue(b, 100)
}

func BenchmarkQueue200(b *testing.B) {
	benchmarkQueue(b, 200)
}

func BenchmarkQueue2(b *testing.B) {
	benchmarkQueue(b, 2)
}

func BenchmarkQueue4(b *testing.B) {
	benchmarkQueue(b, 4)
}

func BenchmarkQueue8(b *testing.B) {
	benchmarkQueue(b, 8)
}

func BenchmarkQueue16(b *testing.B) {
	benchmarkQueue(b, 16)
}

func BenchmarkQueue32(b *testing.B) {
	benchmarkQueue(b, 32)
}

func BenchmarkQueue64(b *testing.B) {
	benchmarkQueue(b, 64)
}

func BenchmarkQueue128(b *testing.B) {
	benchmarkQueue(b, 128)
}

func BenchmarkQueue256(b *testing.B) {
	benchmarkQueue(b, 256)
}

func benchmarkQueue(b *testing.B, n int) {
	internalArraySize = n
	q := New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for _, ok := q.Pop(); ok; _, ok = q.Pop() {
	}
}
