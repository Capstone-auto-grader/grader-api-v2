package graderd

type Service struct {
	schr Scheduler
}

func NewGraderService(schr Scheduler) *Service {
	return &Service{
		schr: schr,
	}
}
