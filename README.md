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
