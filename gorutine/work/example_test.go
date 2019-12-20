package work_test

import (
	"github.com/abxd39/myproject/gorutine/work"
	"log"
	"testing"
)

type example struct {
	index int
}

func (e *example) Run() {
	log.Printf("goroutine index=%v", e.index)

}

func ExampleNewPool() {
	examplePool := work.NewPool(10)
	for i := 0; i <= 10; i++ {
		e := &example{
			index: 1,
		}
		examplePool.Add(e)
		examplePool.Run(10)

	}

}
func TestNewPool(t *testing.T) {
	ExampleNewPool()
}
