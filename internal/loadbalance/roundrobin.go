package loadbalance

func (s *ServerPool) NextIndex() int {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.current++
	if s.current >= len(s.backends) {
		s.current = 0
	}
	return s.current
}

func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()

	for i := 0; i < len(s.backends); i++ {
		idx := (next + i) % len(s.backends)

		if s.backends[idx].IsAlive() {
			if i != 0 {
				s.mux.Lock()
				s.current = idx
				s.mux.Unlock()
			}
			return s.backends[idx]
		}
	}
	return nil
}
