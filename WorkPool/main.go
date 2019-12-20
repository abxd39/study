package main

import (
	"context"
	"fmt"
	"github.com/abxd39/myproject/WorkPool/pool"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	n := runtime.GOMAXPROCS(10)
	log.Println(n)
}

const Max = 100000


func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*1)
	Working(ctx)
	maxgoroutine := 0
	mingoroutine := 0
flag:
	for {
		select {
		case <-ctx.Done():
			break flag
		default:
			gon := runtime.NumGoroutine()
			if maxgoroutine < gon {
				maxgoroutine = gon
				mingoroutine = gon
				log.Printf("goroutine %v", gon)
			}

			if gon < mingoroutine {
				mingoroutine = gon
				log.Printf("mini goroutine %v", gon)
			}
			_ = mingoroutine

		}

	}

	<-ctx.Done()
	time.Sleep(time.Second)
	fmt.Println()
}

func Working(ctx context.Context) {
	dis:=pool.NewDispatcher(4)
	dis.Run(ctx)
	go func() {
		for i := 0; i < 100000; i++ {
			pool.JobQueue1 <- &Payload1{
				name: fmt.Sprintf("palyload1 %v", i),
			}
		}
	}()

	go func() {
		for i := 100000; i < 200000; i++ {
			pool.JobQueue1 <- &Task1{
				Name: fmt.Sprintf("task1 %v", i),
			}
		}
	}()
	go func() {
		for i := 200000; i < 300000; i++ {
			pool.JobQueue1 <- &Payload{
				Name: fmt.Sprintf("payload %v", i),
			}
		}
	}()

	go func() {
		for i := 300000; i < 400000; i++ {
			pool.JobQueue1 <- &Task2{
				Name: fmt.Sprintf("task2 %v", i),
			}
		}
	}()
	go func() {
		for i := 400000; i < 500000; i++ {
			pool.JobQueue1 <- &Task3{
				Name: fmt.Sprintf("task3 %v", i),
			}
		}
	}()
	go func() {
		for i := 500000; i < 600000; i++ {
			pool.JobQueue1 <- &Task4{
				Name: fmt.Sprintf("task4 %v", i),
			}
		}
	}()
	go func() {
		for i := 600000; i < 700000; i++ {
			pool.JobQueue1 <- &Task5{
				Name: fmt.Sprintf("task5 %v", i),
			}
		}
	}()
	go func() {
		for i := 700000; i < 800000; i++ {
			pool.JobQueue1 <- &Task5{
				Name: fmt.Sprintf("task5 %v", i),
			}
		}
	}()
	go func() {
		for i := 800000; i < 900000; i++ {
			pool.JobQueue1 <- &Task6{
				Name: fmt.Sprintf("task6 %v", i),
			}
		}
	}()
	go func() {
		for i := 90000; i < 110000; i++ {
			pool.JobQueue1 <- &Task7{
				Name: fmt.Sprintf("task7 %v", i),
			}
		}
	}()
}

type Payload1 struct {
	name string
}

type Task1 struct {
	Name string
}

type Task2 struct {
	Name string
}
type Task3 struct {
	Name string
}

type Task4 struct {
	Name string
}
type Task5 struct {
	Name string
}
type Task6 struct {
	Name string
}

type Task7 struct {
	Name string
}

func (p *Payload1) Fuck() error {
	//log.Println(p.name)
	//time.Sleep(time.Nanosecond * 2)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	return nil
}

func (t *Task1) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

func (t *Task2) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

func (t *Task3) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

func (t *Task4) Fuck() error {
	//log.Println(t.Name)
	//time.Sleep(time.Nanosecond * 2)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	return nil
}

func (t *Task5) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

func (t *Task6) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

func (t *Task7) Fuck() error {
	//log.Println(t.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}

type Payload struct {
	Name string
}

func (p *Payload) Fuck() error {
	//log.Println(p.Name)
	sum := 0
	for i := 0; i <= Max; i++ {
		i++
		sum += i
	}
	//time.Sleep(time.Nanosecond * 2)
	return nil
}
