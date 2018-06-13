package main

func main() {
	request := HealthReportRequest{
		expectedStatusCode: 200,
		timeOutInSeconds:   5,
		url:                "http://www.yusufaytas.com",
	}
	if reportHealth(request) {
		println("Healthy")
	}
}
