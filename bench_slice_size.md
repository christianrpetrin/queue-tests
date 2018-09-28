# Bench-slice-size
Below benchmark tests probe [queueimpl7](queueimpl7/queueimpl7.go) queue implementation with different internal slice sizes.

## Tests
Format
- Benchmark*N: benchmark the [queueimpl7](queueimpl7/queueimpl7.go) with a internal slice size of N.

Example Tests
- Benchmark*1: benchmark the [queueimpl7](queueimpl7/queueimpl7.go) with a internal slice size of 1.
- Benchmark*10: benchmark the [queueimpl7](queueimpl7/queueimpl7.go) with a internal slice size of 10.
- Benchmark*100: benchmark the [queueimpl7](queueimpl7/queueimpl7.go) with a internal slice size of 100.
<br/>

There are two distinct tests:
- BenchmarkMaxFirstSliceSize: benchmark the size of the first created slice in the queue.
- BenchmarkMaxSubsequentSliceSize: benchmark the size of all subsequent created slices in the queue.

## Results
Results from the [benchmark tests](queueimpl7/benchmark_test.go).

First Slice Size
<br/>
![First Slice Size Performance Results](images/first-slice-size-perf.jpg?raw=true "Benchmark tests")
![First Slice Size Memory Results](images/first-slice-size-mem.jpg?raw=true "Benchmark tests")

Subsequent Slices Size
![Subsequent Slices Size Performance Results](images/slice-size-perf.jpg?raw=true "Benchmark tests")
![Subsequent Slices Size Memory Results](images/slice-size-mem.jpg?raw=true "Benchmark tests")
<br/>


[Curated Results](https://docs.google.com/spreadsheets/d/e/2PACX-1vRnCm7v51Eo5nq66NsGi8aQI6gL14XYJWqaeRJ78ZIWq1pRCtEZfsLD2FcI-gIpUhhTPnkzqDte_SDB/pubhtml?gid=403129988&single=true)
<br/>


### Raw Results

```
go version
go version go1.11 darwin/amd64

sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

go test -benchmem -bench=. -run=^$
goos: darwin
goarch: amd64
pkg: github.com/christianrpetrin/queue-tests/queueimpl7
BenchmarkMaxFirstSliceSize/1-4      	   50000	     25388 ns/op	   24687 B/op	    1017 allocs/op
BenchmarkMaxFirstSliceSize/2-4      	   50000	     25617 ns/op	   24719 B/op	    1018 allocs/op
BenchmarkMaxFirstSliceSize/4-4      	   50000	     25380 ns/op	   24783 B/op	    1019 allocs/op
BenchmarkMaxFirstSliceSize/8-4      	   50000	     26824 ns/op	   24911 B/op	    1020 allocs/op
BenchmarkMaxFirstSliceSize/16-4     	   50000	     25719 ns/op	   25167 B/op	    1021 allocs/op
BenchmarkMaxFirstSliceSize/32-4     	   50000	     26050 ns/op	   25679 B/op	    1022 allocs/op
BenchmarkMaxFirstSliceSize/64-4     	   50000	     26476 ns/op	   26703 B/op	    1023 allocs/op
BenchmarkMaxFirstSliceSize/128-4    	   50000	     27369 ns/op	   26671 B/op	    1022 allocs/op
BenchmarkMaxSubsequentSliceSize/1-4 	   20000	     76836 ns/op	   53967 B/op	    2752 allocs/op
BenchmarkMaxSubsequentSliceSize/10-4         	   50000	     33602 ns/op	   29007 B/op	    1184 allocs/op
BenchmarkMaxSubsequentSliceSize/20-4         	   50000	     30537 ns/op	   27599 B/op	    1096 allocs/op
BenchmarkMaxSubsequentSliceSize/30-4         	   50000	     28628 ns/op	   27471 B/op	    1068 allocs/op
BenchmarkMaxSubsequentSliceSize/40-4         	   50000	     27703 ns/op	   26895 B/op	    1052 allocs/op
BenchmarkMaxSubsequentSliceSize/50-4         	   50000	     27517 ns/op	   28815 B/op	    1044 allocs/op
BenchmarkMaxSubsequentSliceSize/60-4         	   50000	     28714 ns/op	   27951 B/op	    1038 allocs/op
BenchmarkMaxSubsequentSliceSize/70-4         	   50000	     27543 ns/op	   27503 B/op	    1034 allocs/op
BenchmarkMaxSubsequentSliceSize/80-4         	   50000	     26641 ns/op	   26543 B/op	    1030 allocs/op
BenchmarkMaxSubsequentSliceSize/90-4         	   50000	     27078 ns/op	   27791 B/op	    1028 allocs/op
BenchmarkMaxSubsequentSliceSize/100-4        	   50000	     27361 ns/op	   28527 B/op	    1026 allocs/op
BenchmarkMaxSubsequentSliceSize/200-4        	   50000	     25956 ns/op	   28271 B/op	    1018 allocs/op
BenchmarkMaxSubsequentSliceSize/2-4          	   30000	     59811 ns/op	   40015 B/op	    1880 allocs/op
BenchmarkMaxSubsequentSliceSize/4-4          	   30000	     42925 ns/op	   33039 B/op	    1444 allocs/op
BenchmarkMaxSubsequentSliceSize/8-4          	   50000	     36946 ns/op	   29551 B/op	    1226 allocs/op
BenchmarkMaxSubsequentSliceSize/16-4         	   50000	     30597 ns/op	   27951 B/op	    1118 allocs/op
BenchmarkMaxSubsequentSliceSize/32-4         	   50000	     28273 ns/op	   27343 B/op	    1064 allocs/op
BenchmarkMaxSubsequentSliceSize/64-4         	   50000	     26969 ns/op	   26895 B/op	    1036 allocs/op
BenchmarkMaxSubsequentSliceSize/128-4        	   50000	     27316 ns/op	   26671 B/op	    1022 allocs/op
BenchmarkMaxSubsequentSliceSize/256-4        	   50000	     26221 ns/op	   28623 B/op	    1016 allocs/op
BenchmarkMaxSubsequentSliceSize/512-4        	   50000	     25882 ns/op	   28559 B/op	    1012 allocs/op
BenchmarkMaxSubsequentSliceSize/1024-4       	   50000	     25674 ns/op	   28527 B/op	    1010 allocs/op
PASS
ok  	github.com/christianrpetrin/queue-tests/queueimpl7	51.410s
```
