package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/renzoruiz98/go-load-balancer/internal/loadbalance"
	"github.com/renzoruiz98/go-load-balancer/internal/proxy"
)

/*
In general term the function main run the server and verify the status the server
*/

func main() {
	serverList := []string{
		"http://localhost:3001",
		"http://localhost:3002",
		"http://localhost:3003",
	}

	serverPool := loadbalance.ServerPool{}

	for _, target := range serverList {
		parsedURL, err := url.Parse(target)
		if err != nil {
			log.Fatal(err)
		}

		rp, _ := proxy.NewProxy(target)

		backend := &loadbalance.Backend{
			URL:          parsedURL,
			Alive:        true,
			ReverseProxy: rp,
		}
		serverPool.AddBackend(backend)
		fmt.Printf("Backend register %s\n", target)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		peer := serverPool.GetNextPeer()
		if peer != nil {
			peer.ReverseProxy.ServeHTTP(w, r)
			return
		}
		http.Error(w, "No backends available", http.StatusServiceUnavailable)
	})

	port := "8080"
	fmt.Printf("lb listening on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
