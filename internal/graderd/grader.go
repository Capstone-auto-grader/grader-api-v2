package graderd

type Service struct {
	schr    Scheduler
	db      Database
	webAddr string
	maxJobs int
}

func NewGraderService(schr Scheduler, db Database, webAddr string, maxJobs int) *Service {
	return &Service{
		schr:    schr,
		db:      db,
		webAddr: webAddr,
		maxJobs: maxJobs,
	}
}
