package graderd

import (
	docker_client "github.com/Capstone-auto-grader/grader-api-v2/internal/docker-client"
	syncmap "github.com/Capstone-auto-grader/grader-api-v2/internal/sync-map"
)

type Service struct {
	client docker_client.DockerClient
	mp      *syncmap.SyncMap
	jobChan chan<- string
	webAddr string
}

func NewGraderService(client docker_client.DockerClient, mp *syncmap.SyncMap, jobChan chan<- string, webAddr string) *Service {
	return &Service{
		client:    client,
		mp: 	 mp,
		jobChan: jobChan,
		webAddr: webAddr,
	}
}
