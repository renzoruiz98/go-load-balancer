package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renzoruiz98/go-load-balancer/internal/loadbalance"
	"github.com/renzoruiz98/go-load-balancer/internal/proxy"
	"github.com/renzoruiz98/go-load-balancer/internal/telemetry"
)

/*
In general term the function main run the server and verify the status the server
*/

func main() {
	serverList := []string{
		"http://backend1:80",
		"http://backend2:80",
		"http://backend3:80",
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
		start := time.Now()

		peer := serverPool.GetNextPeer()
		if peer != nil {
			backendName := peer.URL.Host
			peer.ReverseProxy.ServeHTTP(w, r)
			duration := time.Since(start).Seconds()
			telemetry.RequestDuration.WithLabelValues(backendName).Observe(duration)
			telemetry.TotalRequests.WithLabelValues(backendName, "200").Inc()
			return
		}
		telemetry.TotalRequests.WithLabelValues("unknown", "503").Inc()
		http.Error(w, "servers not available", http.StatusServiceUnavailable)
	})

	serverPool.StartHealthCheck(10 * time.Second)
	port := "8080"
	fmt.Printf("lb listening on port %s\n", port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
