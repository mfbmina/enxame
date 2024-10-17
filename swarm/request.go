package swarm

import (
	"net/http"
	"time"
)

type HTTPResponse struct {
	StatusCode int           `json:"status_code"`
	Time       time.Duration `json:"time"`
	Path       string        `json:"path"`
}

func Swarm(path, method string, requestsPerUser, concurrentUsers int) ([]HTTPResponse, error) {
	totalRequests := requestsPerUser * concurrentUsers
	channel := make(chan HTTPResponse, totalRequests)
	responses := make([]HTTPResponse, totalRequests)

	for i := 0; i < concurrentUsers; i++ {
		req, _ := http.NewRequest(method, path, nil)
		go userRequests(req, requestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		responses[i] = <-channel
	}

	return responses, nil
}

func userRequests(req *http.Request, requestsPerUser int, channel chan HTTPResponse) {
	for i := 0; i < requestsPerUser; i++ {
		startTime := time.Now()

		client := &http.Client{}
		resp, err := client.Do(req)

		elapsedTime := time.Since(startTime) / time.Millisecond

		if err != nil {
			channel <- HTTPResponse{StatusCode: 0, Path: req.URL.String(), Time: elapsedTime}
			continue
		}

		channel <- HTTPResponse{StatusCode: resp.StatusCode, Path: req.URL.String(), Time: elapsedTime}
	}
}
