package lib

import (
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Time       time.Duration
	Path       string
}

func Swarm(path string, requestsPerUser int, concurrentUsers int) {
	totalRequests := requestsPerUser * concurrentUsers
	channel := make(chan Response, totalRequests)

	fmt.Println("Calling path", path, "...")
	for i := 0; i < concurrentUsers; i++ {
		go userRequests(path, requestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		resp := <-channel
		fmt.Println(resp.StatusCode, resp.Time, resp.Path)
	}

	fmt.Println("Done!")
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
