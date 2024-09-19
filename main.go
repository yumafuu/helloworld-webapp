package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	// This is list of TCP connection that checked at the start of the application
	// Env: TCP_CONNECTION_ADDRS
	// Required: false
	// Default: []string{}
	tcpAddrs = []string{}

	// This is the timeout for TCP connection check
	// format: int
	// Env: TCP_CONNECTION_TIMEOUT
	// Default: 5
	tcpTimeout = 5

	// This is the port that the application will listen to
	// Env: PORT
	// Default: 8080
	port = ""
)

func init() {
	tcpAddrs = strings.Split(os.Getenv("TCP_CONNECTION_ADDRS"), ",")
	timeoutStr := os.Getenv("TCP_CONNECTION_TIMEOUT")
	if timeoutStr != "" {
		var err error
		tcpTimeout, err = strconv.Atoi(timeoutStr)
		if err != nil {
			log.Fatalf("Failed to convert timeout to int: %v", err)
		}
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("TCP Connections: ", tcpAddrs)
	fmt.Println("TCP Timeout: ", tcpTimeout)
	fmt.Println("Port: ", port)
}

func printRequestDetail(r *http.Request) {
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		printRequestDetail(r)
		fmt.Fprintf(w, "Hello, World!")
	})

	for _, addr := range tcpAddrs {
		conn, err := net.DialTimeout("tcp", addr, time.Duration(tcpTimeout)*time.Second)

		if err != nil {
			log.Fatalf("Failed to connect to %s: %v", addr, err)
		}
		log.Printf("Connected to %s\n", addr)
		defer conn.Close()
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
