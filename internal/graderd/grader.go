package graderd

type Service struct {
	schr    Scheduler
	db      Database
	webAddr string
}

func NewGraderService(schr Scheduler, webAddr string) *Service {
	return &Service{
		schr:    schr,
		webAddr: webAddr,
	}
}
