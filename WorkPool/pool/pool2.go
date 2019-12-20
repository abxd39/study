package pool

import (
	"context"
)

var JobQueue1 chan Job = make(chan Job, 10)

//所有的传入池中的结构必须实现一下接口
type Job interface {
	Fuck() error
}

type Task struct {
	WorkPool   chan chan Job
	JobChannel chan Job
}

func newWorker(workerPool chan chan Job) Task {
	chanJob := make(chan Job, 100)
	workerPool <- chanJob
	return Task{
		WorkPool:   workerPool,
		JobChannel: chanJob,
		//quit:       make(chan bool),
	}
}

func (t Task) start(ctx context.Context) {
	go func() {
		for {
			select {
			case job := <-t.JobChannel:
				job.Fuck()
			case <-ctx.Done():
				close(t.JobChannel)
				return
			}
		}
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run(ctx context.Context) {
	for i := 0; i < d.maxWorkers; i++ {
		worker := newWorker(d.WorkerPool)
		worker.start(ctx)
	}
	go d.dispatch(ctx)
}

func (d *Dispatcher) dispatch(ctx context.Context) {
	chanJob := <-d.WorkerPool
BreakFlag:
	for {

		select {
		case job := <-JobQueue1:
			chanJob <- job
		case <-ctx.Done():
			close(JobQueue1)
			close(d.WorkerPool)
			break BreakFlag
		}
	}
}
