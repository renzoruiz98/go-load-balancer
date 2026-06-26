package loadbalance

import (
	"log"
	"net"
	"net/url"
	"time"
)

/*
stablish connection tcp with the URL
*/
func isBackendLive(u *url.URL) bool {
	timeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		log.Printf("Error connecting to backend: %s\n", u.Host)
		return false
	}
	_ = conn.Close()
	return true
}
