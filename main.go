package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Unable to get IP address"
}

func greet(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	currentTime := time.Now().Format(time.RFC1123)

	fmt.Fprintf(w, "Hostname: %s, IP: %s, Time: %s", hostname, getIP(), currentTime)
}

func main() {
	http.HandleFunc("/", greet)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Printf("Server running at http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
