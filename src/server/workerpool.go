package server

type WorkerPool struct {
	JobQueue chan Job
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	pool := &WorkerPool{
		JobQueue: make(chan Job, maxWorkers),
	}

	for i := 0; i < maxWorkers; i++ {
		go pool.worker()
	}

	return pool
}

func (p *WorkerPool) worker() {
	for job := range p.JobQueue {
		handleConnection(job.Conn)
	}
}
