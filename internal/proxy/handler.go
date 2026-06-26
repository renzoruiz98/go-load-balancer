package proxy

import (
	"net/http/httputil"
	"net/url"
)

/*
func newproxy applied one inverse proxy and return one proxy configurate or invalid URL
*/

func NewProxy(target string) (*httputil.ReverseProxy, error) {
	parsedURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(parsedURL), nil
}
