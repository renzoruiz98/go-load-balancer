package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/renzoruiz98/go-load-balancer/internal/proxy"
)

/*
In general term the function main run the server and verify the status the server
*/

func main() {
	targetURL := "http://localhost:3001"

	reverseProxy, err := proxy.NewProxy(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	port := "8080"
	fmt.Printf("lb listening on port %s\n", port)
	fmt.Println("Redirecting all traffic to -> %s\\n\"", targetURL)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
