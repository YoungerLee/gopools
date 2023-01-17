package pools

import (
	"encoding/json"
	"testing"
)

type PoolId uint64

func init() {
	Register[PoolId, Student](poolId)
}

const poolId PoolId = 1

type Student struct {
	ID     uint64  `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender int     `json:"genger"`
	Class  int32   `json:"class"`
	Marks  []int32 `json:"marks"`
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := Get[PoolId, Student](poolId)
		v.ID = 123456
		v.Name = "Anonymous"
		v.Age = 22
		v.Gender = 1
		v.Class = 10
		v.Marks = []int32{95, 96, 97, 8, 99, 100}
		b, _ := json.Marshal(v)
		Put(poolId, v)
		v1 := Get[PoolId, Student](poolId)
		json.Unmarshal(b, v1)
		Put(poolId, v1)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := new(Student)
		v.ID = 123456
		v.Name = "Anonymous"
		v.Age = 22
		v.Gender = 1
		v.Class = 10
		v.Marks = []int32{95, 96, 97, 8, 99, 100}
		b, _ := json.Marshal(v)
		v1 := new(Student)
		json.Unmarshal(b, v1)
	}
}
