# go-pools
A memory pool manager implemented with generics.

## Installation
```shell
go get github.com/YoungerLee/go-pools
```

## Quickstart
Memory pools must be registered before they can be used. Type parameters of pool id type and pooled object type should be provided when registering.
```go
package main

import (
	"fmt"

	"github.com/YoungerLee/go-pools"
)

type PoolId uint64

type Student struct {
	ID     uint64  `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender int     `json:"genger"`
	Class  int32   `json:"class"`
	Marks  []int32 `json:"marks"`
}

const STUDENT_POOL_ID PoolId = 1

func init() {
	pools.Register[PoolId, Student](STUDENT_POOL_ID)
}

func main() {
	s := pools.Get[PoolId, Student](STUDENT_POOL_ID)
	s.ID = 123456
	s.Name = "Anonymous"
	s.Age = 22
	s.Gender = 1
	s.Class = 10
	s.Marks = []int32{95, 96, 97, 8, 99, 100}
	fmt.Printf("%p|%+v", s, s)
	pools.Put(STUDENT_POOL_ID, s)
}

```
## Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/YoungerLee/go-pools
BenchmarkWithPool-16              289014              4538 ns/op             392 B/op         16 allocs/op
BenchmarkWithoutPool-16           237301              5084 ns/op             640 B/op         22 allocs/op
```
