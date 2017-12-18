package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/crash", crashAppHandler)
	http.HandleFunc("/showvars", showVarsHandler)
	go waitForStoppage()
	fmt.Println("Starting the server.")
	http.ListenAndServe(":5000", nil)
}

func helloWorldHandler(resp http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	io.WriteString(resp, fmt.Sprintf("Hello world from %s - app version: %s", hostname, "v4"))
}

func crashAppHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Commiting suicide. Bye-bye")
	os.Exit(1)
}

func showVarsHandler(resp http.ResponseWriter, req *http.Request) {
	u, _ := os.LookupEnv("username_secret")
	s, _ := os.LookupEnv("password_secret")
	io.WriteString(resp, fmt.Sprintf("Username: %s Secret: %s", u, s))
}

func waitForStoppage() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Stopping... Got signal:", <-signalChan)
	os.Exit(0)
}
