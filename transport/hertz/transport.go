package hertz

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-kratos/kratos/v2/errors"
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

func DedaultServerResponseEncoder(ctx context.Context, c *app.RequestContext, res any) {
	// data, err := sonic.MarshalString(res)
	// if err != nil {
	// 	c.String(consts.StatusInternalServerError, err.Error())
	// 	return
	// }
	switch v := res.(type) {
	// case proto.Message:
	// 	c.ProtoBuf(consts.StatusOK, v)
	case string:
		c.String(consts.StatusOK, v)
	default:
		c.JSON(consts.StatusOK, res)
	}
}

func DedaultErrorEncoder(ctx context.Context, c *app.RequestContext, err error) {
	se := errors.FromError(err)
	body, err := sonic.MarshalString(se)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	c.Response.Header.Set("Content-Type", "application/json")
	c.String(int(se.Code), body)
}
