package test

import (
	"net"
	"testing"

	"github.com/kanengo/ktp/transport"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := transport.GetLocalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}

func TestSplitHostPort(t *testing.T) {
	host, port, err := net.SplitHostPort(":8888")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(host, port)
}
