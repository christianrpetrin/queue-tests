# Queue-tests [![Build Status](https://travis-ci.com/christianrpetrin/queue-tests.svg?branch=master)](https://travis-ci.com/christianrpetrin/queue-tests) 
Package queue-tests tests a number of queue implementations for performance and memory consumption.

## How to Use
This package is mostly useful to run its benchmark tests. Run below commands to clone and get the repo ready for running the tests locally.

```
git clone https://github.com/christianrpetrin/queue-tests.git
go get ./...
```

The benchmark tests are full lifecycle tests which includes, in each test iteration:
- Queue initialization
- Adding N to values the queue
- Removing N values from the queue

The tests probe adding and removing below number of items to the queues.

- 0
- 1
- 10
- 100
- 1000   // 1k
- 10000  // 10k
- 100000 // 100k

The 0 test tests initialization only.

Some tests probe the queues performance by adding 3 items and then removing 1 instead of adding all items and them removing all items.
To keep the runtime short, these tests will only execute for 100 and 10.000 items.

To run the tests locally, from the repo root directory, execute below command:

```
go test -benchmem -bench=. -run=^$
```

## Tests
See [bench_queue.md](bench_queue.md) and [bench_slice_size.md](bench_slice_size.md) for details.


## Supported Go Versions
The bench tests on this package automatically compile and run for below Go versions.

- "1.7.x"
- "1.8.x"
- "1.9.x"
- "1.10.x"
- "1.11.x"

## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"
