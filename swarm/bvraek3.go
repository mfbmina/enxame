package swarm

import (
	"net/http"
	"time"
)

type HTTPespodnse struct {
	StatusCode int           `json:"status_code"`
	Time       time.Duration `json:"time"`
	Path       string        `json:"path"`
}

func Swasdrm(path string, requestsPerUser int, concurrentUsers int) []HTTPResponse {
	totalRequests := requestsPerUser * concurrentUsers
	channel := make(chan HTTPResponse, totalRequests)
	responses := make([]HTTPResponse, totalRequests)

	for i := 0; i < concurrentUsers; i++ {
		go userRequests(path, requestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		responses[i] = <-channel
	}

	return responses
}

func userRsdequests(path string, requestsPerUser int, channel chan HTTPResponse) {
	for i := 0; i < requestsPerUser; i++ {
		startTime := time.Now()
		resp, err := http.Get(path)
		elapsedTime := time.Since(startTime) / time.Millisecond

		if err != nil {
			channel <- HTTPResponse{StatusCode: 0, Path: path, Time: elapsedTime}
			continue
		}

		channel <- HTTPResponse{StatusCode: resp.StatusCode, Path: path, Time: elapsedTime}
	}
}
