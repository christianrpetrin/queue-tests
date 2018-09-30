# Bench-Queue
Below benchmark tests probe some of the most promising open source queue implementations as well as a few experimental implementations in this package. It also compares the queue implementations with using the [standard list package](https://github.com/golang/go/tree/master/src/container/list) as well as using buffered channels as queues.

## Tests
There are two sets of distinct benchmark tests:
1) Adds N values to the queue sequentially and them removes them all afterwards
2) Adds N (N=3) values to the queue, removes 1 item; removes all remaining elements in the queue afterwards

The tests probe some of the most promising and fastest unbounded open source queue implementations, namely [gammazero](https://github.com/gammazero/deque), [phf](https://github.com/phf/go-queue) and [juju](https://github.com/juju/utils/tree/master/deque). These queue implementations were chosen due to their design, high quality implementation and very good performance.

The tests also probe a number of experimental queue implementations. Benchmark*Impl1 test a queue implementation that was
suggested [here](https://stackoverflow.com/a/26863706). Benchmark*Impl2 is the same as Benchmark*Impl1, but tests a slightly different implementation where instead of removing the first element from the queue, it uses pointers to point to the next element. Benchmark*Impl3, 4, 5, 6 and 7 uses linked slices as its underlying data structure testing different implementations.

The tests:
- BenchmarkList: benchmark the standard list package using it as an unbounded FIFO queue. This is a linked list based queue implementation.
- BenchmarkChannel: benchmark buffered channels using it as a bounded FIFO queue.
- BenchmarkGammazero: benchmark the [gammazero](https://github.com/gammazero/deque) package using it as an unbounded FIFO queue. This is a slice, ring based queue implementation.
- BenchmarkPhf: benchmark the [phf](https://github.com/phf/go-queue) package using it as an unbounded FIFO queue. This is a slice, ring based queue implementation as well.
- BenchmarkJuju: benchmark the [juju](https://github.com/juju/utils/tree/master/deque) package using it as an unbounded FIFO queue. This is a linked list based queue implementation that uses slices as linked list values.
- BenchmarkImpl1: benchmark a custom queue implementation that stores the values in a simple slice. Pop removes the first slice element. This is a slice based implementation that tests [this](https://stackoverflow.com/a/26863706) suggestion.
- BenchmarkImpl2: benchmark a custom queue implementation that stores the values in a simple slice. Pop moves the current position to next one instead of removing the first element. This is a slice based implementation similarly to "Benchmark*Impl1", but differs in the fact that it uses pointers to point to the current first element in the queue instead of removing the first element.
- BenchmarkImpl3: benchmark a custom queue implementation that stores the values in linked slices. This implementation tests the queue performance when controlling the length and current positions in the slices using the builtin len and append functions.
- BenchmarkImpl4: benchmark a custom queue implementation that stores the values in linked arrays. This implementation tests the queue performance when controlling the length and current positions in the arrays using simple local variables instead of the builtin len and append functions.
- BenchmarkImpl5: benchmark a custom queue implementation that stores the values in linked slices. This implementation tests the queue performance when storing the "next" pointer as part of the values slice instead of having it as a separate "next" field. The next element is stored in the last position of the internal slice, which is a reserved position.
- BenchmarkImpl6: benchmark a custom queue implementation that stores the values in linked slices. This implementation tests the queue performance when performing lazy creation of the first slice as well as starting with a slice of size 1 and doubling up to 128.
- BenchmarkImpl7: benchmark a custom queue implementation that stores the values in linked slices. This implementation tests the queue performance when performing lazy creation of the internal slice as well as starting with a 1-sized slice, allowing it to grow up to 16 by using the builtin append function. Subsequent slices are created with 128 fixed size.
<br/>


## Results
Below a few interesting results.
<br/>

Initialization time only<br/>
Performance<br/>
![ns/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-0-items-perf.jpg?raw=true "Benchmark tests")
<br/>

Memory<br/>
![B/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-0-items-mem.jpg?raw=true "Benchmark tests")

Add and remove 1k items<br/>
Performance<br/>
![ns/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-1k-items-perf.jpg?raw=true "Benchmark tests")

Memory<br/>
![B/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-1k-items-mem.jpg?raw=true "Benchmark tests")
<br/>

Add and remove 100k items<br/>
Performance<br/>
![ns/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-100k-items-perf.jpg?raw=true "Benchmark tests")

Memory<br/>
![B/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-100k-items-mem.jpg?raw=true "Benchmark tests")
<br/>

Aggregated Results<br/>
Performance<br/>
![ns/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-line-perf.jpg?raw=true "Benchmark tests")

Memory<br/>
![B/op](https://github.com/christianrpetrin/queue-tests/blob/master/images/queue-line-mem.jpg?raw=true "Benchmark tests")
<br/>

Detailed, curated results can be found [here](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=668319604&single=true)

Aggregated, curated results can be found [here](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=582031751&single=true)
<br/>

Benchmark tests can be found [here](benchmark_test.go).


### Raw Results
```
go version
go version go1.11 darwin/amd64

sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

go test -benchmem -bench=. -run=^$
goos: darwin
goarch: amd64
pkg: github.com/christianrpetrin/queue-tests
BenchmarkList/0-4      	50000000	        35.2 ns/op	      48 B/op	       1 allocs/op
BenchmarkList/1-4      	20000000	        77.3 ns/op	      96 B/op	       2 allocs/op
BenchmarkList/10-4     	 3000000	       572 ns/op	     600 B/op	      20 allocs/op
BenchmarkList/100-4    	  200000	      5910 ns/op	    5640 B/op	     200 allocs/op
BenchmarkList/1000-4   	   30000	     54784 ns/op	   56040 B/op	    2000 allocs/op
BenchmarkList/10000-4  	    2000	    633409 ns/op	  560040 B/op	   20000 allocs/op
BenchmarkList/100000-4 	     100	  13630519 ns/op	 5600048 B/op	  200000 allocs/op
BenchmarkChannel/0-4   	50000000	        30.4 ns/op	      96 B/op	       1 allocs/op
BenchmarkChannel/1-4   	20000000	        87.6 ns/op	     112 B/op	       1 allocs/op
BenchmarkChannel/10-4  	 2000000	       702 ns/op	     248 B/op	      10 allocs/op
BenchmarkChannel/100-4 	  200000	      6805 ns/op	    1688 B/op	     100 allocs/op
BenchmarkChannel/1000-4         	   20000	     66862 ns/op	   16184 B/op	    1000 allocs/op
BenchmarkChannel/10000-4        	    2000	    671592 ns/op	  161912 B/op	   10000 allocs/op
BenchmarkChannel/100000-4       	     200	   6744860 ns/op	 1602810 B/op	  100000 allocs/op
BenchmarkGammazero/0-4          	2000000000	         1.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGammazero/1-4          	20000000	       102 ns/op	     256 B/op	       1 allocs/op
BenchmarkGammazero/10-4         	 5000000	       334 ns/op	     328 B/op	      10 allocs/op
BenchmarkGammazero/100-4        	  300000	      4005 ns/op	    6424 B/op	     106 allocs/op
BenchmarkGammazero/1000-4       	   50000	     33787 ns/op	   56632 B/op	    1012 allocs/op
BenchmarkGammazero/10000-4      	    5000	    322555 ns/op	  472696 B/op	   10018 allocs/op
BenchmarkGammazero/100000-4     	     100	  11897764 ns/op	 7090944 B/op	  100026 allocs/op
BenchmarkPhf/0-4                	50000000	        27.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkPhf/1-4                	30000000	        42.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkPhf/10-4               	 2000000	       667 ns/op	     696 B/op	      15 allocs/op
BenchmarkPhf/100-4              	  300000	      4497 ns/op	    6792 B/op	     111 allocs/op
BenchmarkPhf/1000-4             	   50000	     34876 ns/op	   57000 B/op	    1017 allocs/op
BenchmarkPhf/10000-4            	    5000	    338644 ns/op	  473064 B/op	   10023 allocs/op
BenchmarkPhf/100000-4           	     100	  11953773 ns/op	 7091312 B/op	  100031 allocs/op
BenchmarkJuju/0-4               	 5000000	       349 ns/op	    1184 B/op	       4 allocs/op
BenchmarkJuju/1-4               	 5000000	       365 ns/op	    1184 B/op	       4 allocs/op
BenchmarkJuju/10-4              	 2000000	       611 ns/op	    1256 B/op	      13 allocs/op
BenchmarkJuju/100-4             	  500000	      3603 ns/op	    4184 B/op	     109 allocs/op
BenchmarkJuju/1000-4            	   50000	     31376 ns/op	   26840 B/op	    1051 allocs/op
BenchmarkJuju/10000-4           	    5000	    317426 ns/op	  253400 B/op	   10471 allocs/op
BenchmarkJuju/100000-4          	     300	   4088804 ns/op	 2525629 B/op	  104689 allocs/op
BenchmarkImpl1/0-4              	200000000	         6.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkImpl1/1-4              	30000000	        46.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkImpl1/10-4             	 3000000	       507 ns/op	     568 B/op	      14 allocs/op
BenchmarkImpl1/100-4            	  500000	      3120 ns/op	    4360 B/op	     108 allocs/op
BenchmarkImpl1/1000-4           	   50000	     24177 ns/op	   40744 B/op	    1010 allocs/op
BenchmarkImpl1/10000-4          	    5000	    314997 ns/op	  746344 B/op	   10022 allocs/op
BenchmarkImpl1/100000-4         	     100	  13300508 ns/op	10047344 B/op	  100029 allocs/op
BenchmarkImpl2/0-4              	200000000	         6.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkImpl2/1-4              	30000000	        45.7 ns/op	      16 B/op	       1 allocs/op
BenchmarkImpl2/10-4             	 3000000	       512 ns/op	     568 B/op	      14 allocs/op
BenchmarkImpl2/100-4            	  500000	      3132 ns/op	    4872 B/op	     107 allocs/op
BenchmarkImpl2/1000-4           	   50000	     24041 ns/op	   40744 B/op	    1010 allocs/op
BenchmarkImpl2/10000-4          	    5000	    383255 ns/op	  905961 B/op	   10019 allocs/op
BenchmarkImpl2/100000-4         	     100	  14913362 ns/op	10047344 B/op	  100029 allocs/op
BenchmarkImpl3/0-4              	 3000000	       443 ns/op	    2080 B/op	       2 allocs/op
BenchmarkImpl3/1-4              	 3000000	       459 ns/op	    2080 B/op	       2 allocs/op
BenchmarkImpl3/10-4             	 2000000	       658 ns/op	    2152 B/op	      11 allocs/op
BenchmarkImpl3/100-4            	  500000	      2375 ns/op	    2872 B/op	     101 allocs/op
BenchmarkImpl3/1000-4           	  100000	     22384 ns/op	   24632 B/op	    1015 allocs/op
BenchmarkImpl3/10000-4          	   10000	    231449 ns/op	  244312 B/op	   10157 allocs/op
BenchmarkImpl3/100000-4         	     500	   2890319 ns/op	 2426557 B/op	  101563 allocs/op
BenchmarkImpl4/0-4              	 5000000	       331 ns/op	    2304 B/op	       1 allocs/op
BenchmarkImpl4/1-4              	 5000000	       335 ns/op	    2304 B/op	       1 allocs/op
BenchmarkImpl4/10-4             	 3000000	       520 ns/op	    2376 B/op	      10 allocs/op
BenchmarkImpl4/100-4            	 1000000	      2194 ns/op	    3096 B/op	     100 allocs/op
BenchmarkImpl4/1000-4           	  100000	     20907 ns/op	   26424 B/op	    1007 allocs/op
BenchmarkImpl4/10000-4          	   10000	    215338 ns/op	  262008 B/op	   10078 allocs/op
BenchmarkImpl4/100000-4         	     500	   2734631 ns/op	 2601725 B/op	  100781 allocs/op
BenchmarkImpl5/0-4              	 3000000	       410 ns/op	    2048 B/op	       1 allocs/op
BenchmarkImpl5/1-4              	 3000000	       425 ns/op	    2048 B/op	       1 allocs/op
BenchmarkImpl5/10-4             	 2000000	       611 ns/op	    2120 B/op	      10 allocs/op
BenchmarkImpl5/100-4            	 1000000	      2259 ns/op	    2840 B/op	     100 allocs/op
BenchmarkImpl5/1000-4           	  100000	     22009 ns/op	   24600 B/op	    1014 allocs/op
BenchmarkImpl5/10000-4          	   10000	    227122 ns/op	  244280 B/op	   10156 allocs/op
BenchmarkImpl5/100000-4         	     200	   6381480 ns/op	 2439005 B/op	  101574 allocs/op
BenchmarkImpl6/0-4              	1000000000	         2.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkImpl6/1-4              	20000000	        71.7 ns/op	      48 B/op	       2 allocs/op
BenchmarkImpl6/10-4             	 3000000	       551 ns/op	     440 B/op	      17 allocs/op
BenchmarkImpl6/100-4            	  500000	      3167 ns/op	    3048 B/op	     113 allocs/op
BenchmarkImpl6/1000-4           	   50000	     25199 ns/op	   24808 B/op	    1027 allocs/op
BenchmarkImpl6/10000-4          	    5000	    253783 ns/op	  244488 B/op	   10169 allocs/op
BenchmarkImpl6/100000-4         	     500	   2954444 ns/op	 2426733 B/op	  101575 allocs/op
BenchmarkImpl7/0-4              	2000000000	         1.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkImpl7/1-4              	20000000	        68.0 ns/op	      48 B/op	       2 allocs/op
BenchmarkImpl7/10-4             	 3000000	       579 ns/op	     600 B/op	      15 allocs/op
BenchmarkImpl7/100-4            	  500000	      3098 ns/op	    3400 B/op	     107 allocs/op
BenchmarkImpl7/1000-4           	   50000	     25366 ns/op	   25160 B/op	    1021 allocs/op
BenchmarkImpl7/10000-4          	    5000	    262077 ns/op	  242760 B/op	   10161 allocs/op
BenchmarkImpl7/100000-4         	     500	   2991916 ns/op	 2427085 B/op	  101569 allocs/op
PASS
ok  	github.com/christianrpetrin/queue-tests	152.800s
```
