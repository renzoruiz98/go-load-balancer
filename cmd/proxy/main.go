package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
In general term the function main run the server and verify the status the server
*/

func main() {
	port := "8080"
	fmt.Printf("initializing lb in port %s\n", port)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))

		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalf("Error in server", err)
		}
	})

}
