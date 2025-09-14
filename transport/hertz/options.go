package hertz

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

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

func WithResponseEncoder(encoder func(ctx context.Context, c *app.RequestContext, res any)) ServerOption {
	return func(s *Server) {
		s.ResponseEncoder = encoder
	}
}

func WithErrorEncoder(encoder func(ctx context.Context, c *app.RequestContext, err error)) ServerOption {
	return func(s *Server) {
		s.ErrorEncoder = encoder
	}
}
