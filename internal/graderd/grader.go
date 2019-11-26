package graderd

type Service struct {
	schr    Scheduler
	db      Database
	webAddr string
	results chan *Task
}

func NewGraderService(schr Scheduler, db Database, webAddr string, maxJobs int) *Service {
	return &Service{
		schr:    schr,
		db:      db,
		webAddr: webAddr,
		results: make(chan *Task, maxJobs),
	}
}
