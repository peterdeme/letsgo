package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Starting the server.")
	http.ListenAndServe(":5000", nil)
}

func helloWorldHandler(resp http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	io.WriteString(resp, fmt.Sprintf("Hello world from %s - app version: %s", hostname, "v1"))
}
