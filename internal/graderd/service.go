package graderd

type Service struct {
	schd 	Scheduler

	webAddr string
}

func NewGraderService(schd Scheduler, webAddr string) *Service {
	return &Service{
		schd:    schd,

		webAddr: webAddr,
	}
}
