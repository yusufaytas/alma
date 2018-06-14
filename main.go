package main

import "time"

func main() {
	request := HealthCheckRequest{
		expectedStatusCode: 200,
		timeOutInSeconds:   5,
		url:                "http://www.yusufaytas.com",
	}

	for {
		time.Sleep(2 * time.Second)
		go CheckHealth(request, []HealthListener{ConsoleListener{}})
	}
}
