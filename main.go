package main

import (
	"fmt"
	"net/http"
)

func main() {
	ping("http://www.yusufaytas.com")
}

func ping(domain string) {
	response, _ := http.Get(domain)
	fmt.Printf(response.Status)
}
