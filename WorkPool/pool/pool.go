package pool

import "sync"

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

type Worker interface {
	Run() error
}

func NewPool(maxQueueSize int) *Pool {
	p := Pool{
		work: make(chan Worker, maxQueueSize),
	}
	p.run(maxQueueSize)
	return &p
}

func (p *Pool) Add(w Worker) {
	p.wg.Add(1)
	p.work <- w
}

func (p *Pool) run(MaxPushWorker int) {
	for i := 0; i < MaxPushWorker; i++ {
		go func() {
			for w := range p.work {
				w.Run()
				p.wg.Done()
			}
		}()
	}
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
