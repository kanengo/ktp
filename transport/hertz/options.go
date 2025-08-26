package hertz

import "time"

type ServerOption func(*Server)

func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.addr = addr
	}
}

func WithExitWaitTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.exitWaitTimeout = timeout
	}
}
