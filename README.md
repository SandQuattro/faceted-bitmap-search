# faceted-bitmap-search

## Performance difference between bitmap search and goroutine execution, which do nothing...
```
pkg: go-bitmask-search/searcher
Benchmark_Search
Start searching in users total set: 100000000
time elapsed: 418.317583ms
total users found 99987114
....
Start searching in users total set: 100000000
time elapsed: 181.679417ms
total users found 74992476

Start searching in users total set: 100000000
time elapsed: 184.135125ms
total users found 74992476

Benchmark_Search-14    	       6	 192256667 ns/op	762521510 B/op	      11 allocs/op
PASS
```

```
sending 99998268 notifications to users
notifications processing elapsed: 33.001048042s
total notifications: 99998268
Exiting.
```

# TODO
- [ ] add parallel chunked dataset processing

Inspired by https://github.com/1F47E/faceted-bitmap