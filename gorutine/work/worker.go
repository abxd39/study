package work

import "sync"

type Worker interface {
	Run()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewPool(maxQueueSize int) *Pool {
	p := Pool{
		work: make(chan Worker, maxQueueSize),
	}
	return &p
}

func (p *Pool) Add(w Worker) {
	p.wg.Add(1)
	p.work <- w
}

func (p *Pool) Run(maxWorker int) {
	for i := 0; i < maxWorker; i++ {
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