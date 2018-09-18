# Bench-Queue
Below benchmark tests probe some of the most promising open source queue implementations as well as a few other implementations in this package. It also compares the queue implementations with using the standard list and ring packages as well as using channels as queues.


## Results

Results from benchmark add tests. Below tests only add items to the queue.

![Results from benchmark add tests](images/queue_add_tests.jpg?raw=true "Add tests")

[Spreadsheet](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=0&single=true)

Results from benchmark add and remove tests. Below tests add items to the queue and remove them all afterwards as part of the test.

![Results from benchmark add and remove tests](images/queue_add_remove_tests.jpg?raw=true "Add and remove tests")

[Spreadsheet](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=2037026674&single=true)

[Benchmark tests](benchmark_test.go)


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
BenchmarkAddList-4              	30000000	       154 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       153 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       140 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       182 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       172 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       208 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       195 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       180 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       172 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddList-4              	30000000	       160 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRing-4              	100000000	       149 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       114 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       101 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRing-4              	50000000	       124 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRing-4              	50000000	        82.1 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       110 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRing-4              	50000000	        96.7 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       101 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       104 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRing-4              	50000000	       103 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddChannel-4           	200000000	        26.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.9 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        28.1 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.8 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.1 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        26.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        27.1 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddChannel-4           	200000000	        28.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddGammazero-4         	50000000	        82.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        82.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        81.0 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	       101 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        77.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        86.6 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        82.6 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        85.7 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        80.2 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddGammazero-4         	50000000	        79.6 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	100000000	        91.3 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        70.5 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        75.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        72.2 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        74.2 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        75.3 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        78.5 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        91.8 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        74.4 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddPhf-4               	50000000	        92.0 ns/op	      50 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        46.6 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        45.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        41.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        54.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        38.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        46.8 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        46.3 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        54.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        37.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddJuju-4              	100000000	        59.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        30.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        33.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        31.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        30.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        31.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        30.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        32.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	200000000	        31.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        30.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl1-4             	100000000	        35.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        30.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        30.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        31.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        34.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        30.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	200000000	        31.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        32.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        37.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        31.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl2-4             	100000000	        35.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        31.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        30.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        32.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        30.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	100000000	        34.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        35.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        31.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        31.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        29.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl3-4             	200000000	        30.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        34.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        38.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        36.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        36.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        40.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        35.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        34.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        34.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        36.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl4-4             	100000000	        43.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       230 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       190 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       154 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       162 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       198 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	30000000	       171 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	30000000	       165 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       183 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       176 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl5-4             	20000000	       187 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       219 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       176 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       195 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       171 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       187 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       169 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       175 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	30000000	       179 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	20000000	       151 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddImpl6-4             	30000000	       174 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddRemoveList-4        	30000000	       176 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       178 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       172 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       171 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       188 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       138 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       170 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       169 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       247 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveList-4        	30000000	       155 ns/op	      56 B/op	       2 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       112 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       118 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        99.8 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        99.4 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       106 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        96.9 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       100 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	        95.1 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       100 ns/op	      40 B/op	       2 allocs/op
BenchmarkAddRemoveRing-4        	50000000	       102 ns/op	      40 B/op	       1 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        55.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.7 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.7 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        53.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.4 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        51.5 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveChannel-4     	100000000	        52.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       108 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       121 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       112 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       108 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       110 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       131 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       106 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       135 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       114 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveGammazero-4   	50000000	       124 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       121 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       121 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       127 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       118 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       127 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       129 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       108 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       119 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       105 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemovePhf-4         	50000000	       110 ns/op	      72 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        53.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        47.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        59.7 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        48.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        47.5 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        65.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        50.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        45.9 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        67.0 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveJuju-4        	100000000	        66.4 ns/op	      25 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        63.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        32.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        72.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        33.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        34.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        37.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        33.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl1-4       	100000000	        31.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        64.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        31.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        40.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        73.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        33.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        32.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        67.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        32.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        67.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl2-4       	100000000	        63.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        33.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        33.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        66.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        36.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        68.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        69.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        32.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl3-4       	100000000	        67.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        40.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        55.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        40.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        42.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        38.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        41.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        39.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        41.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        38.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl4-4       	100000000	        44.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       201 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       203 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       205 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       195 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       206 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       198 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       201 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       180 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       190 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl5-4       	20000000	       194 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	30000000	       197 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       196 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	30000000	       200 ns/op	      92 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       190 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       215 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       191 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       188 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       191 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       198 ns/op	      88 B/op	       1 allocs/op
BenchmarkAddRemoveImpl6-4       	20000000	       190 ns/op	      88 B/op	       1 allocs/op
PASS
ok  	github.com/christianrpetrin/queue-tests	1339.481s
```
