package pool_test

import (
	"context"
	"fmt"
	"github.com/abxd39/myproject/WorkPool/pool"
	"testing"
)

type Example struct {
}

func (e *Example) Run() error {
	fmt.Println("example")
	return nil
}

var p *pool.Pool

func ExampleNewPool() {
	p =pool.NewPool(10)
	p.Add(&Example{})
}

func BenchmarkNewDispatcher(b *testing.B) {
	d := pool.NewDispatcher(10)
	d.Run(context.TODO())
	for i := 0; i < b.N; i++ {
		pool.JobQueue1 <- &pool.Payload{
			Name: "benchmem",
		}
	}
}

func ExampleNewDispatcher() {
	d := pool.NewDispatcher(10)
	d.Run(context.TODO())
	//生产这队列
	pool.JobQueue1 <- &pool.Payload{
		Name: "benchmem",
	}

}
