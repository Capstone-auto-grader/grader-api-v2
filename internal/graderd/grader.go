package graderd

type Service struct {
	schr    Scheduler
	db      Database
	webAddr string
}

func NewGraderService(schr Scheduler, db Database, webAddr string) *Service {
	return &Service{
		schr:    schr,
		db:      db,
		webAddr: webAddr,
	}
}
