# match
Byte matching in Go

Under development....

# speed

Match a 4 byte string in 1024 characters. There is 1 match.
```
BenchmarkMatch4String-8          5000000               292 ns/op        3495.75 MB/s
```

Match a 8 bytes in 32K byte slice. There is 1 match.
```
BenchmarkMatch8-8                 100000             11821 ns/op        2771.86 MB/s
```

Match a 8 bytes in 32K byte slice, as well as matching first 4 bytes. There is 1 match.
```
BenchmarkMatch8And4-8             100000             14423 ns/op        2271.88 MB/s
```

Determine if two 256 byte slice matches. All bytes match.
```
BenchmarkMatchLen256-8          30000000                46.7 ns/op      5485.93 MB/s
```
