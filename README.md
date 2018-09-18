# Queue-tests [![Build Status](https://travis-ci.com/christianrpetrin/queue-tests.svg?branch=master)](https://travis-ci.com/christianrpetrin/queue-tests) 
Package queue-tests tests a number of queue implementations for performance and memory consumption.

## How to Use
This package is mostly useful to run its benchmark tests. Run below commands to clone and run the tests locally.

```
git clone https://github.com/christianrpetrin/queue-tests.git
go get ./...
go test -bench=. -benchmem
```


## Tests
See [bench_queue.md](bench_queue.md) and [bench_array_size.md](bench_array_size.md) for details.

Tests
- BenchmarkList*: benchmark the standard list package using it as a FIFO queue.
- BenchmarkRing*: benchmark the standard ring package using it as a FIFO queue.
- BenchmarkChannel*: benchmark a standard channel using it as a FIFO queue.
- BenchmarkGammazeroDeque*: benchmark the [gammazero](https://github.com/gammazero/deque) package using it as a FIFO queue.
- BenchmarkPhfQueue*: benchmark the [phf](https://github.com/phf/go-queue) package using it as a FIFO queue.
- BenchmarkJujuDeque*: benchmark the [juju](https://github.com/juju/utils/tree/master/deque) package using it as a FIFO queue.
- BenchmarkQueueImpl1*: benchmark a custom queue implementation which uses a linked array with append and len builtin functions.
- BenchmarkQueueImpl2*: benchmark a custom queue implementation which uses a linked array with simple local variables instead of relying on builtin len and append functions.
- BenchmarkQueueImpl3*: benchmark a custom queue implementation which uses a linked array with simple local variables instead of relying on builtin len and append functions. This implementation doesn't accept nil values and returns nil on empty queue. Otherwise it's the same implementation as BenchmarkQueueImpl2*.
- BenchmarkQueueImpl4*: benchmark a custom queue implementation that stores the "next" pointer as part of the values array instead of using a separate field (as in previous custom implementations).
- BenchmarkQueueImpl5*: benchmark a custom queue implementation that stores the values in a simple slice. Pop moves the current position to next one instead of removing the first element from the slice.
- BenchmarkQueueImpl6*: benchmark a custom queue implementation that stores the values in a simple slice. Pop removes the first element from the slice.


## Supported Go Versions
The bench tests on this package automatically compile and run for below Go versions.

- "1.x"
- "1.2.x"
- "1.3.x"
- "1.4.x"
- "1.5.x"
- "1.9.x"
- "1.7.x"
- "1.8.x"
- "1.9.x"
- "1.10.x"
- "1.11.x"


## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"
