package main

import (
	"fmt"
	"net/http"
	"time"
)

type HealthCheckRequest struct {
	url                string
	timeOutInSeconds   int64
	expectedStatusCode int
}

type HealthListener interface {
	onStart(request HealthCheckRequest)
	onSuccess(request HealthCheckRequest)
	onFailure(request HealthCheckRequest)
}

func CheckHealth(request HealthCheckRequest, healthListeners []HealthListener) {
	client := getClient(request.timeOutInSeconds)
	response, err := client.Get(request.url)
	for _, healthListener := range healthListeners {
		healthListener.onStart(request)
	}
	if nil != err || response.StatusCode != request.expectedStatusCode {
		for _, healthListener := range healthListeners {
			healthListener.onFailure(request)
		}
		return
	}
	for _, healthListener := range healthListeners {
		healthListener.onSuccess(request)
	}
}

func getClient(timeoutInSeconds int64) http.Client {
	timeout := time.Duration(timeoutInSeconds * int64(time.Second))
	return http.Client{Timeout: timeout}
}

type ConsoleListener struct {
	startTime time.Time
	endTime   time.Time
}

func (ConsoleListener) onStart(request HealthCheckRequest) {
	fmt.Println("Started connecting to " + request.url)
}

func (ConsoleListener) onSuccess(request HealthCheckRequest) {
	fmt.Println("Successfully connected to " + request.url)
}

func (ConsoleListener) onFailure(request HealthCheckRequest) {
	fmt.Println("Failed to connect to " + request.url)
}
