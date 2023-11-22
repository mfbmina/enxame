package swarm

import (
	"net/http"
	"time"
)

type Response struct {
	StatusCode int           `json:"status_code"`
	Time       time.Duration `json:"time"`
	Path       string        `json:"path"`
}

func Swarm(path string, requestsPerUser int, concurrentUsers int) []Response {
	totalRequests := requestsPerUser * concurrentUsers
	channel := make(chan Response, totalRequests)
	responses := make([]Response, totalRequests)

	for i := 0; i < concurrentUsers; i++ {
		go userRequests(path, requestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		responses[i] = <-channel
	}

	return responses
}

func userRequests(path string, requestsPerUser int, channel chan Response) {
	for i := 0; i < requestsPerUser; i++ {
		startTime := time.Now()
		resp, err := http.Get(path)
		elapsedTime := time.Since(startTime) / time.Millisecond

		if err != nil {
			channel <- Response{StatusCode: 0, Path: path, Time: elapsedTime}
			continue
		}

		channel <- Response{StatusCode: resp.StatusCode, Path: path, Time: elapsedTime}
	}
}
