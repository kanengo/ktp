package hertz

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	kttransport "github.com/go-kratos/kratos/v2/transport"
	"github.com/kanengo/ktp/transport"
)

type Server struct {
	*server.Hertz

	addr            string
	exitWaitTimeout time.Duration
	endpoint        *url.URL
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		addr:            ":8888",
		exitWaitTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(srv)
	}

	srv.init(opts...)

	return srv
}

func (s *Server) init(opts ...ServerOption) {
	for _, opt := range opts {
		opt(s)
	}
	s.Hertz = server.New(
		server.WithHostPorts(s.addr),
		server.WithExitWaitTime(s.exitWaitTimeout),
	)
}

func (s *Server) listenAndEndpoint() error {
	if s.endpoint == nil {
		host, port, err := net.SplitHostPort(s.addr)
		if err != nil {
			return err
		}

		if host == "" {
			ip, _ := transport.GetLocalIP()
			host = ip
		}

		addr := host + ":" + fmt.Sprint(port)
		s.endpoint = &url.URL{
			Scheme: KindHertz,
			Host:   addr,
		}
	}

	return nil
}

func (s *Server) Endpoint() (*url.URL, error) {
	if err := s.listenAndEndpoint(); err != nil {
		return nil, err
	}
	return s.endpoint, nil
}

// Start implements kttransport.Server.
func (s *Server) Start(context.Context) error {
	if err := s.listenAndEndpoint(); err != nil {
		return err
	}
	return s.Run()
}

// Stop implements kttransport.Server.
func (s *Server) Stop(ctx context.Context) error {
	err := s.Hertz.Shutdown(ctx)

	return err
}

var (
	_ kttransport.Server = (*Server)(nil)
)
