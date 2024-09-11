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

## Concurrent chunked dataset processing
```
Created Bitmask: 00010101 00000000 00000000 00000000 
2024/09/11 23:04:26 time
2024/09/11 23:04:26 Start searching in users subset from 21428571 to 28571428, chunk: 3 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 92857141 to 100000000, chunk: 13 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 28571428 to 35714285, chunk: 4 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 35714285 to 42857142, chunk: 5 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 49999999 to 57142856, chunk: 7 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 57142856 to 64285713, chunk: 8 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 64285713 to 71428570, chunk: 9 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 71428570 to 78571427, chunk: 10 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 78571427 to 85714284, chunk: 11 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 42857142 to 49999999, chunk: 6 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 85714284 to 92857141, chunk: 12 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 7142857 to 14285714, chunk: 1 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 0 to 7142857, chunk: 0 using bitmask:                            10101
2024/09/11 23:04:26 Start searching in users subset from 14285714 to 21428571, chunk: 2 using bitmask:                            10101
2024/09/11 23:04:26 time elapsed: 8.425ms
total users found 892718

2024/09/11 23:04:26 time elapsed: 8.435625ms
total users found 895104

2024/09/11 23:04:26 time elapsed: 8.462042ms
total users found 893431

2024/09/11 23:04:26 time elapsed: 8.447167ms
total users found 893429

2024/09/11 23:04:26 time elapsed: 8.676459ms
total users found 893066

2024/09/11 23:04:26 time elapsed: 8.854458ms
total users found 892104

2024/09/11 23:04:26 time elapsed: 8.376333ms
total users found 893396

2024/09/11 23:04:26 time elapsed: 8.825708ms
total users found 893035

2024/09/11 23:04:26 time elapsed: 8.946709ms
total users found 892170

2024/09/11 23:04:26 time elapsed: 10.027ms
total users found 892073

2024/09/11 23:04:26 time elapsed: 11.079708ms
total users found 892091

2024/09/11 23:04:26 time elapsed: 11.548833ms
total users found 893112

2024/09/11 23:04:26 time elapsed: 10.856958ms
total users found 892726

2024/09/11 23:04:26 time elapsed: 11.329375ms
total users found 892932

2024/09/11 23:04:26 Found 12501387 options in 100000000 users, time: 29.758416ms
2024/09/11 23:04:26 sending 12501387 notifications to users
2024/09/11 23:04:29 notifications processing elapsed: 3.834681291s
total notifications: 12501387

```

# TODO
- [x] add parallel chunked dataset processing
- [ ] add channel data processing

Inspired by https://github.com/1F47E/faceted-bitmap