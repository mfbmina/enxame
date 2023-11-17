package lib

import (
	"fmt"
	"net/http"
)

func Swarm(path string, requestsPerUser int, concurrentUsers int) {
	totalRequests := requestsPerUser * concurrentUsers
	channel := make(chan http.Response, totalRequests)

	fmt.Println("Calling path", path, "...")
	for i := 0; i < concurrentUsers; i++ {
		go userRequests(path, requestsPerUser, channel)
	}

	for i := 0; i < totalRequests; i++ {
		resp := <-channel
		fmt.Println(resp.StatusCode)
	}

	fmt.Println("Done!")
}

func userRequests(path string, requestsPerUser int, channel chan http.Response) {
	for i := 0; i < requestsPerUser; i++ {
		resp, err := http.Get(path)
		if err != nil {
			channel <- http.Response{StatusCode: 0}
			continue
		}

		channel <- *resp
	}
}
