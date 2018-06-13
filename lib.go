package main

import (
	"net/http"
	"time"
)

type HealthReportRequest struct {
	url                string
	timeOutInSeconds   int64
	expectedStatusCode int
}

func reportHealth(request HealthReportRequest) bool {
	client := getClient(request.timeOutInSeconds)
	response, err := client.Get(request.url)
	if nil != err {
		return false
	}
	return response.StatusCode == request.expectedStatusCode
}

func getClient(timeoutInSeconds int64) http.Client {
	timeout := time.Duration(timeoutInSeconds * int64(time.Second))
	return http.Client{Timeout: timeout}
}
