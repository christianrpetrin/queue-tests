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
	"strconv"
	"testing"

	"github.com/christianrpetrin/queue-tests/queueimpl1"
	"github.com/christianrpetrin/queue-tests/queueimpl2"
	"github.com/christianrpetrin/queue-tests/queueimpl3"
	"github.com/christianrpetrin/queue-tests/queueimpl4"
	"github.com/christianrpetrin/queue-tests/queueimpl5"
	"github.com/christianrpetrin/queue-tests/queueimpl6"
	"github.com/christianrpetrin/queue-tests/queueimpl7"
	gammazero "github.com/gammazero/deque"
	juju "github.com/juju/utils/deque"
	phf "github.com/phf/go-queue/queue"
)

type testData struct {
	count  int
	remove bool
}

var (
	tests = []testData{
		{count: 0},
		{count: 1},
		{count: 10},
		{count: 100, remove: true},
		{count: 1000},                // 1k
		{count: 10000, remove: true}, //10k
		{count: 100000},              // 100k
	}

	// Used to store temp values, avoiding any compiler optimizations.
	tmp  interface{}
	tmp2 bool
)

func BenchmarkList(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				l := list.New()

				for i := 0; i < test.count; i++ {
					l.PushBack(i)

					if test.remove && i > 0 && i%3 == 0 {
						l.Remove(l.Front())
					}
				}
				for e := l.Front(); e != nil; e = e.Next() {
					tmp = e.Value
				}
			}
		})
	}
}

func BenchmarkChannel(b *testing.B) {
	for i, test := range tests {
		// Only run the first 8 tests for channel as it is bounded and don't support very large channel sizes.
		if i <= 8 {
			b.Run(strconv.Itoa(test.count), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					c := make(chan int, test.count)

					for i := 0; i < test.count; i++ {
						c <- i
					}
					for i := 0; i < test.count; i++ {
						tmp = <-c
					}
				}
			})
		}
	}
}

func BenchmarkGammazero(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				var q gammazero.Deque

				for i := 0; i < test.count; i++ {
					q.PushBack(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp = q.PopFront()
					}
				}
				for q.Len() > 0 {
					tmp = q.PopFront()
				}
			}
		})
	}
}

func BenchmarkPhf(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := phf.New()

				for i := 0; i < test.count; i++ {
					q.PushBack(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp = q.PopFront()
					}
				}
				for q.Len() > 0 {
					tmp = q.PopFront()
				}
			}
		})
	}
}

func BenchmarkJuju(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := juju.New()

				for i := 0; i < test.count; i++ {
					q.PushBack(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.PopFront()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.PopFront()
				}
			}
		})
	}
}

func BenchmarkImpl1(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl1.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}
func BenchmarkImpl2(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl2.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}

func BenchmarkImpl3(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl3.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}

func BenchmarkImpl4(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl4.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}

func BenchmarkImpl5(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl5.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}

func BenchmarkImpl6(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl6.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}

func BenchmarkImpl7(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := queueimpl7.New()

				for i := 0; i < test.count; i++ {
					q.Push(i)

					if test.remove && i > 0 && i%3 == 0 {
						tmp, tmp2 = q.Pop()
					}
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.Pop()
				}
			}
		})
	}
}
