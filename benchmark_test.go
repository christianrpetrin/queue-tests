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

package tests

import (
	"container/list"
	"container/ring"
	"testing"

	"github.com/christianrpetrin/queue-tests/queueimpl1"
	"github.com/christianrpetrin/queue-tests/queueimpl2"
	"github.com/christianrpetrin/queue-tests/queueimpl3"
	"github.com/christianrpetrin/queue-tests/queueimpl4"
	"github.com/christianrpetrin/queue-tests/queueimpl5"
	"github.com/christianrpetrin/queue-tests/queueimpl6"
	gammazero "github.com/gammazero/deque"
	juju "github.com/juju/utils/deque"
	phf "github.com/phf/go-queue/queue"
)

func BenchmarkAddList(b *testing.B) {
	q := list.New()

	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}
}

func BenchmarkAddRing(b *testing.B) {
	r := ring.New(b.N)

	for n := 0; n < b.N; n++ {
		r.Value = n
		r = r.Next()
	}
}

func BenchmarkAddChannel(b *testing.B) {
	c := make(chan int, b.N)

	for n := 0; n < b.N; n++ {
		c <- n
	}
}

func BenchmarkAddGammazero(b *testing.B) {
	var q gammazero.Deque

	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}
}

func BenchmarkAddPhf(b *testing.B) {
	q := phf.New()

	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}
}

func BenchmarkAddJuju(b *testing.B) {
	q := juju.New()

	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}
}

func BenchmarkAddImpl1(b *testing.B) {
	q := queueimpl1.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkAddImpl2(b *testing.B) {
	q := queueimpl2.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkAddImpl3(b *testing.B) {
	q := queueimpl3.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkAddImpl4(b *testing.B) {
	q := queueimpl4.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkAddImpl5(b *testing.B) {
	q := queueimpl5.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

func BenchmarkAddImpl6(b *testing.B) {
	q := queueimpl6.New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

// --------------------------------------------------------------

func BenchmarkAddRemoveList(b *testing.B) {
	l := list.New()
	for n := 0; n < b.N; n++ {
		l.PushBack(n)
	}

	for e := l.Front(); e != nil; e = e.Next() {
	}
}

func BenchmarkAddRemoveRing(b *testing.B) {
	r := ring.New(b.N)
	for n := 0; n < b.N; n++ {
		r.Value = n
		r = r.Next()
	}

	for j := 0; j < b.N; j++ {
		r = r.Next()
	}
}

func BenchmarkAddRemoveChannel(b *testing.B) {
	c := make(chan int, b.N)
	for n := 0; n < b.N; n++ {
		c <- n
	}

	for n := 0; n < b.N; n++ {
		<-c
	}
}

func BenchmarkAddRemoveGammazero(b *testing.B) {
	defer func() {
		// gammazero deque implementation throws panic on empty deuque with no way to find out if the deuque is empty before hand.
		// panic is a normal condition of this deque implementation.
		recover()
	}()
	var q gammazero.Deque
	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}

	for v := q.PopFront(); v != nil; v = q.PopFront() {
	}
}

func BenchmarkAddRemovePhf(b *testing.B) {
	q := phf.New()
	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}

	for e := q.PopFront(); e != nil; e = q.PopFront() {
	}
}

func BenchmarkAddRemoveJuju(b *testing.B) {
	q := juju.New()
	for n := 0; n < b.N; n++ {
		q.PushBack(n)
	}

	for _, ok := q.PopFront(); ok; _, ok = q.PopFront() {
	}
}

func BenchmarkAddRemoveImpl1(b *testing.B) {
	q := queueimpl1.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for _, ok := q.Pop(); ok; _, ok = q.Pop() {
	}
}

func BenchmarkAddRemoveImpl2(b *testing.B) {
	q := queueimpl2.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for _, ok := q.Pop(); ok; _, ok = q.Pop() {
	}
}

func BenchmarkAddRemoveImpl3(b *testing.B) {
	q := queueimpl3.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for v := q.Pop(); v != nil; v = q.Pop() {
	}
}

func BenchmarkAddRemoveImpl4(b *testing.B) {
	q := queueimpl4.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for v := q.Pop(); v != nil; v = q.Pop() {
	}
}

func BenchmarkAddRemoveImpl5(b *testing.B) {
	q := queueimpl5.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for v := q.Pop(); v != nil; v = q.Pop() {
	}
}

func BenchmarkAddRemoveImpl6(b *testing.B) {
	q := queueimpl6.New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for v := q.Pop(); v != nil; v = q.Pop() {
	}
}
