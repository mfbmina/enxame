package lib

import (
	"fmt"
	"net/http"
)

func Service(path string) {
	maxRequestsPerUser := 5
	concurrentUsers := 2
	totalRequests := maxRequestsPerUser * concurrentUsers
	channel := make(chan http.Response, totalRequests)

	fmt.Println("Calling path", path, "...")
	for i := 0; i < concurrentUsers; i++ {
		go doRequest(path, maxRequestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		resp := <-channel
		fmt.Println(resp.StatusCode)
	}

	fmt.Println("Done!")
}

func doRequest(path string, maxRequestsPerUser int, channel chan http.Response) {
	for i := 0; i < maxRequestsPerUser; i++ {
		resp, err := http.Get(path)
		if err != nil {
			channel <- http.Response{StatusCode: 0}
			continue
		}

		channel <- *resp
	}
}
