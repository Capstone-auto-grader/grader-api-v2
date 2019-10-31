package graderd

type Service struct {
	schr    Scheduler
	webAddr string
}

func NewGraderService(schr Scheduler, webAddr string) *Service {
	return &Service{
		schr:    schr,
		webAddr: webAddr,
	}
}
