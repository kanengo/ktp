package hertz

import (
	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KindHertz = "hertz"
)

type Transport struct {
	endpoint string
}

// Endpoint implements transport.Transporter.
func (tr *Transport) Endpoint() string {
	panic("unimplemented")
}

// Kind implements transport.Transporter.
func (tr *Transport) Kind() transport.Kind {
	return KindHertz
}

// Operation implements transport.Transporter.
func (tr *Transport) Operation() string {
	panic("unimplemented")
}

// ReplyHeader implements transport.Transporter.
func (tr *Transport) ReplyHeader() transport.Header {
	panic("unimplemented")
}

// RequestHeader implements transport.Transporter.
func (tr *Transport) RequestHeader() transport.Header {
	panic("unimplemented")
}

var _ transport.Transporter = (*Transport)(nil)
