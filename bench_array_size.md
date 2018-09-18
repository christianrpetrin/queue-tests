# Bench-array-size
Below benchmark tests probe queueimpl1 queue implementation with different internal array sizes.

## Tests
Format
- BenchmarkQueueN: benchmark the [queueimpl1](queueimpl1/queueimpl1.go) with a internal array size of N.

Example Tests
- BenchmarkQueue1: benchmark the [queueimpl1](queueimpl1/queueimpl1.go) with a internal array size of 1.
- BenchmarkQueue10: benchmark the [queueimpl1](queueimpl1/queueimpl1.go) with a internal array size of 10.
- BenchmarkQueue100: benchmark the [queueimpl1](queueimpl1/queueimpl1.go) with a internal array size of 100.
<br/>

## Results

![Results from benchmark tests](images/array_size_tests.jpg?raw=true "Tests")

[Spreadsheet](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=403129988&single=true)

[Benchmark tests](queueimpl1/benchmark_test.go)
<br/>


### Raw Results

```
go version
go version go1.10.3 darwin/amd64

sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

go test -bench=. -benchmem -count 10 -benchtime 5s -timeout 60m
goos: darwin
goarch: amd64
pkg: github.com/christianrpetrin/queue-tests/queueimpl1
BenchmarkQueue1-4     	50000000	       161 ns/op	      56 B/op	       3 allocs/op
BenchmarkQueue1-4     	50000000	       180 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       145 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       112 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       156 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       165 ns/op	      56 B/op	       3 allocs/op
BenchmarkQueue1-4     	50000000	       118 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       116 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       112 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue1-4     	50000000	       145 ns/op	      56 B/op	       2 allocs/op
BenchmarkQueue10-4    	100000000	        86.7 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        63.3 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        82.2 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	200000000	        67.8 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	200000000	        80.2 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        52.7 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        52.1 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	200000000	        47.3 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        85.4 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue10-4    	100000000	        74.9 ns/op	      27 B/op	       1 allocs/op
BenchmarkQueue20-4    	200000000	        64.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        60.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	200000000	        71.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        79.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        71.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	200000000	        79.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        75.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        79.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        76.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue20-4    	100000000	        77.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        71.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        70.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	200000000	        71.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	200000000	        32.6 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        71.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	200000000	        70.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        72.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	200000000	        75.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        70.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue30-4    	100000000	        70.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue40-4    	200000000	        58.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	200000000	        62.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	100000000	        75.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	100000000	        77.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	200000000	        38.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	200000000	        39.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	200000000	        44.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	100000000	        61.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	100000000	        77.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue40-4    	100000000	        66.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        71.7 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        72.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        73.6 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        73.1 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	200000000	        69.6 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        66.9 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	200000000	        71.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	100000000	        66.9 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	200000000	        73.7 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue50-4    	200000000	        37.9 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        64.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	200000000	        69.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        64.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        63.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        65.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	200000000	        64.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        54.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	200000000	        32.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        69.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue60-4    	100000000	        62.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue70-4    	200000000	        31.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        81.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	200000000	        66.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        67.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        64.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	200000000	        32.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        67.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        70.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	200000000	        66.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue70-4    	100000000	        67.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	100000000	        73.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        74.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	100000000	        69.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	100000000	        71.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        36.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        37.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        53.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	100000000	        78.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        81.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue80-4    	200000000	        42.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        41.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        46.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        42.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	100000000	        83.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        38.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        41.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	100000000	        70.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        43.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        40.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue90-4    	200000000	        35.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        74.1 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        69.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        67.7 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        68.5 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	200000000	        63.4 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        69.5 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        67.4 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        69.6 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        67.9 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue100-4   	100000000	        68.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue200-4   	200000000	        61.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        65.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        69.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        60.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        69.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        62.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        68.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	300000000	        33.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	100000000	        63.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue200-4   	200000000	        63.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       102 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	100000000	       111 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       119 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	100000000	       105 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       111 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       141 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       133 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	100000000	       112 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       102 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue2-4     	50000000	       109 ns/op	      40 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        87.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        75.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        75.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	       108 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        80.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	       109 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	       107 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        99.3 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        84.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue4-4     	100000000	        99.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        91.3 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        51.6 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        53.3 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	200000000	        92.4 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        93.4 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        56.8 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        93.7 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        77.9 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        50.1 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue8-4     	100000000	        97.3 ns/op	      28 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        92.8 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        79.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        80.3 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        81.5 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        79.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        81.6 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        75.8 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	100000000	        74.0 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	200000000	        73.7 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue16-4    	200000000	        76.3 ns/op	      26 B/op	       1 allocs/op
BenchmarkQueue32-4    	200000000	        52.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	100000000	        64.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	200000000	        70.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	100000000	        72.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	200000000	        61.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	100000000	        75.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	100000000	        67.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	200000000	        69.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	200000000	        34.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue32-4    	100000000	        70.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        33.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        66.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        33.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        32.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	100000000	        70.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        34.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        34.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        66.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	100000000	        72.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue64-4    	200000000	        65.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        31.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	100000000	        64.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	100000000	        63.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        34.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        30.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        62.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        30.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        30.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	200000000	        31.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue128-4   	300000000	        38.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        35.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        33.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        53.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        42.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	100000000	        84.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        38.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	100000000	        77.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        71.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	100000000	        64.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkQueue256-4   	200000000	        59.8 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	github.com/christianrpetrin/queue-tests/queueimpl1	2060.564s
```
