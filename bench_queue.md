# Bench-Queue
Below benchmark tests probe some of the most promising open source queue implementations as well as a few other implementations in this package. It also compares the queue implementations with using the standard list and ring packages as well as using channels as queues.

## Tests
- Benchmark*List: benchmark the standard list package using it as a FIFO queue.
- Benchmark*Ring: benchmark the standard ring package using it as a FIFO queue.
- Benchmark*Channel: benchmark a standard channel using it as a FIFO queue.
- Benchmark*Gammazero: benchmark the [gammazero](https://github.com/gammazero/deque) package using it as a FIFO queue.
- Benchmark*Phf: benchmark the [phf](https://github.com/phf/go-queue) package using it as a FIFO queue.
- Benchmark*Juju: benchmark the [juju](https://github.com/juju/utils/tree/master/deque) package using it as a FIFO queue.
- Benchmark*Impl1: benchmark a custom queue implementation which uses a linked array with append and len builtin functions.
- Benchmark*Impl2: benchmark a custom queue implementation which uses a linked array with simple local variables instead of relying on builtin len and append functions.
- Benchmark*Impl3: benchmark a custom queue implementation which uses a linked array with simple local variables instead of relying on builtin len and append functions. This implementation doesn't accept nil values and returns nil on empty queue. Otherwise it's the same implementation as Benchmark*Impl2.
- Benchmark*Impl4: benchmark a custom queue implementation that stores the "next" pointer as part of the values array instead of using a separate field (as in previous custom implementations).
- Benchmark*Impl5: benchmark a custom queue implementation that stores the values in a simple slice. Pop moves the current position to next one instead of removing the first element from the slice.
- Benchmark*Impl6: benchmark a custom queue implementation that stores the values in a simple slice. Pop removes the first element from the slice.
<br/>


## Results
Results from benchmark add tests. Below tests only add items to the queue.

![Results from benchmark add tests](images/queue_add_tests.jpg?raw=true "Add tests")

[Spreadsheet](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=0&single=true)

Results from benchmark add and remove tests. Below tests add items to the queue and remove them all afterwards as part of the test.

![Results from benchmark add and remove tests](images/queue_add_remove_tests.jpg?raw=true "Add and remove tests")

[Spreadsheet](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=2037026674&single=true)

[Benchmark tests](benchmark_test.go)
<br/>

### Raw Results

```
go version
go version go1.10.3 darwin/amd64

sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

go test -bench=. -benchmem -count 10 -benchtime 3s -timeout 60m
goos: darwin
goarch: amd64
pkg: github.com/christianrpetrin/queue-tests
BenchmarkAddList-4              	30000000	       197 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       177 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       119 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       163 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       154 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       191 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	20000000	       158 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	20000000	       168 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       163 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       169 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRing-4              	50000000	        99.3 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	        91.6 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	30000000	       105 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       115 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	30000000	       100 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       111 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	        93.9 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	        96.2 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       100 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       105 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddChannel-4           	200000000	        27.9 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.9 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.1 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.5 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.5 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddGammazero-4         	50000000	        83.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        90.1 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	       105 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        87.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        99.6 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        97.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        88.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        96.4 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	       104 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        79.0 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        85.2 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        78.2 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        85.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        87.1 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	       100 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        97.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        88.9 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        86.5 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        86.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        86.4 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        65.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        44.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        47.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        43.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        52.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        66.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        64.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        60.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        62.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        42.6 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        32.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        34.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        30.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        34.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        28.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        30.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        29.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        30.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        30.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        30.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        31.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        29.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        29.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        31.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        30.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        31.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        30.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        33.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        30.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        30.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        32.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        32.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        32.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        32.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        30.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        38.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        37.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        32.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        39.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        51.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        57.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        37.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        32.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        38.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        40.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       176 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       169 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       175 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       165 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       161 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       158 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       183 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	30000000	       175 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       170 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       159 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       172 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       157 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	30000000	       182 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       167 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       169 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       169 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       164 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       164 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       158 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       158 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveList-4        	30000000	       176 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       170 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	20000000	       217 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       175 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       154 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       174 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       182 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       178 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       170 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       181 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       104 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        98.6 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       115 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       115 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       124 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        87.7 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       126 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       104 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        97.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       117 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.5 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.4 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.4 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.8 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.4 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        53.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.5 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveGammazero-4   	30000000	       107 ns/op	      61 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       127 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       100 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       110 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       125 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       126 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       125 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       127 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       110 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       108 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       116 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       126 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       112 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       125 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       122 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       107 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       108 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	        98.7 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       120 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       128 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        50.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        47.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        65.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        44.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        47.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        50.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        53.2 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        69.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        46.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        59.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        32.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        32.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        64.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        32.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        36.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	50000000	        73.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        32.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        71.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        32.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        66.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        31.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        70.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        37.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        67.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        71.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        68.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        34.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        75.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        34.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        66.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        34.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        31.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	50000000	        73.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        31.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        77.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        59.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        40.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        40.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        42.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        56.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        40.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        39.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        56.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        59.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       198 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       198 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       196 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       203 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       193 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       203 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       201 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       188 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	30000000	       209 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       194 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       207 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       205 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       193 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       205 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       204 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       202 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       201 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       205 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       195 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       214 ns/op	      88 B/op	       1 allocs/op
PASS
ok  	github.com/christianrpetrin/queue-tests	1256.261s
```
