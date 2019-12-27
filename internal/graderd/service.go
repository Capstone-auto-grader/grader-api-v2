package graderd

import (
	syncmap "github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
)

type Service struct {
	schd 	Scheduler
	mp      *syncmap.SyncMap
	jobChan chan<- string
	webAddr string
}

func NewGraderService(schd Scheduler, mp *syncmap.SyncMap, jobChan chan<- string, webAddr string) *Service {
	return &Service{
		schd:    schd,
		mp:      mp,
		jobChan: jobChan,
		webAddr: webAddr,
	}
}
