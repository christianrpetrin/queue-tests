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

package queueimpl1

import (
	"testing"
)

func BenchmarkQueueImpl1AddWith1SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 1)
}

func BenchmarkQueueImpl1AddWith10SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 10)
}

func BenchmarkQueueImpl1AddWith20SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 20)
}

func BenchmarkQueueImpl1AddWith30SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 30)
}

func BenchmarkQueueImpl1AddWith40SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 40)
}

func BenchmarkQueueImpl1AddWith50SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 50)
}

func BenchmarkQueueImpl1AddWith60SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 60)
}

func BenchmarkQueueImpl1AddWith70SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 70)
}

func BenchmarkQueueImpl1AddWith80SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 80)
}

func BenchmarkQueueImpl1AddWith90SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 90)
}

func BenchmarkQueueImpl1AddWith100SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 100)
}

func BenchmarkQueueImpl1AddWith200SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 200)
}

func BenchmarkQueueImpl1AddWith2SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 2)
}

func BenchmarkQueueImpl1AddWith4SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 4)
}

func BenchmarkQueueImpl1AddWith8SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 8)
}

func BenchmarkQueueImpl1AddWith16SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 16)
}

func BenchmarkQueueImpl1AddWith32SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 32)
}

func BenchmarkQueueImpl1AddWith64SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 64)
}

func BenchmarkQueueImpl1AddWith128SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 128)
}

func BenchmarkQueueImpl1AddWith256SizedArray(b *testing.B) {
	benchmarkQueueImpl1Add(b, 256)
}

func benchmarkQueueImpl1Add(b *testing.B, n int) {
	internalArraySize = n
	q := New()

	for n := 0; n < b.N; n++ {
		q.Push(n)
	}
}

// --------------------------------------------------------------

// func BenchmarkQueueImpl1RemoveWith1SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 1)
// }

// func BenchmarkQueueImpl1RemoveWith10SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 10)
// }

// func BenchmarkQueueImpl1RemoveWith20SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 20)
// }

// func BenchmarkQueueImpl1RemoveWith30SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 30)
// }

// func BenchmarkQueueImpl1RemoveWith40SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 40)
// }

// func BenchmarkQueueImpl1RemoveWith50SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 50)
// }

// func BenchmarkQueueImpl1RemoveWith60SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 60)
// }

// func BenchmarkQueueImpl1RemoveWith70SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 70)
// }

// func BenchmarkQueueImpl1RemoveWith80SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 80)
// }

// func BenchmarkQueueImpl1RemoveWith90SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 90)
// }

// func BenchmarkQueueImpl1RemoveWith100SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 100)
// }

// func BenchmarkQueueImpl1RemoveWith200SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 200)
// }

// func BenchmarkQueueImpl1RemoveWith2SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 2)
// }

// func BenchmarkQueueImpl1RemoveWith4SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 4)
// }

// func BenchmarkQueueImpl1RemoveWith8SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 8)
// }

// func BenchmarkQueueImpl1RemoveWith16SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 16)
// }

// func BenchmarkQueueImpl1RemoveWith32SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 32)
// }

// func BenchmarkQueueImpl1RemoveWith64SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 64)
// }

// func BenchmarkQueueImpl1RemoveWith128SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 128)
// }

// func BenchmarkQueueImpl1RemoveWith256SizedArray(b *testing.B) {
// 	benchmarkQueueImpl1Remove(b, 256)
// }

// func benchmarkQueueImpl1Remove(b *testing.B, n int) {
// 	internalArraySize = n
// 	q := New()
// 	for n := 0; n < b.N; n++ {
// 		q.Push(n)
// 	}
// 	b.ResetTimer()

// 	for _, ok := q.Pop(); ok; _, ok = q.Pop() {
// 	}
// }

// --------------------------------------------------------------

func BenchmarkQueueImpl1AddAndRemoveWith1SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 1)
}

func BenchmarkQueueImpl1AddAndRemoveWith10SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 10)
}

func BenchmarkQueueImpl1AddAndRemoveWith20SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 20)
}

func BenchmarkQueueImpl1AddAndRemoveWith30SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 30)
}

func BenchmarkQueueImpl1AddAndRemoveWith40SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 40)
}

func BenchmarkQueueImpl1AddAndRemoveWith50SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 50)
}

func BenchmarkQueueImpl1AddAndRemoveWith60SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 60)
}

func BenchmarkQueueImpl1AddAndRemoveWith70SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 70)
}

func BenchmarkQueueImpl1AddAndRemoveWith80SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 80)
}

func BenchmarkQueueImpl1AddAndRemoveWith90SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 90)
}

func BenchmarkQueueImpl1AddAndRemoveWith100SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 100)
}

func BenchmarkQueueImpl1AddAndRemoveWith200SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 200)
}

func BenchmarkQueueImpl1AddAndRemoveWith2SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 2)
}

func BenchmarkQueueImpl1AddAndRemoveWith4SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 4)
}

func BenchmarkQueueImpl1AddAndRemoveWith8SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 8)
}

func BenchmarkQueueImpl1AddAndRemoveWith16SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 16)
}

func BenchmarkQueueImpl1AddAndRemoveWith32SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 32)
}

func BenchmarkQueueImpl1AddAndRemoveWith64SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 64)
}

func BenchmarkQueueImpl1AddAndRemoveWith128SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 128)
}

func BenchmarkQueueImpl1AddAndRemoveWith256SizedArray(b *testing.B) {
	benchmarkQueueImpl1AddAndRemove(b, 256)
}

func benchmarkQueueImpl1AddAndRemove(b *testing.B, n int) {
	internalArraySize = n
	q := New()
	for n := 0; n < b.N; n++ {
		q.Push(n)
	}

	for _, ok := q.Pop(); ok; _, ok = q.Pop() {
	}
}
