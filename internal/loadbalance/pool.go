package loadbalance

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

// destination server
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

/*
ServerPool manage the collectio the backend and container the lb
sync.Mutex for protected the round robin
*/
type ServerPool struct {
	backends []*Backend
	current  int
	mux      sync.Mutex
}

/*
@AddBackend add new server in secure pool
@SetAlive update the state server blocking temp the memory
@IsAlive read the state the server
*/

func (s *ServerPool) AddBackend(b *Backend) {
	s.backends = append(s.backends, b)
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	alive := b.Alive
	b.mux.RUnlock()
	return alive
}
