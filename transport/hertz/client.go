package hertz

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/kanengo/ktp/transport"
	"github.com/kanengo/ku/unsafex"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	*client.Client
}

func NewClient(opts ...config.ClientOption) *Client {
	cli, _ := client.NewClient(opts...)
	c := &Client{cli}

	return c
}

func (c *Client) Do(ctx context.Context, uri string, method string, in, out any, opts ...config.RequestOption) error {
	req, res := &protocol.Request{}, &protocol.Response{}
	req.SetOptions(opts...)
	req.SetMethod(method)
	req.SetRequestURI(uri)

	err := c.RequestEncoder(ctx, req, in)
	if err != nil {
		return err
	}

	err = c.Client.Do(ctx, req, res)
	if err != nil {
		return err
	}

	err = c.ResponseDecoder(ctx, res, out)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RequestEncoder(ctx context.Context, req *protocol.Request, in any) error {
	var body []byte
	var err error

	ct := "application/json;charset=utf-8"

	switch v := in.(type) {
	case proto.Message:
		body, err = proto.Marshal(v)
		ct = "application/x-protobuf;charset=utf-8"
	case string:
		body = []byte(v)
		ct = "text/plain;charset=utf-8"
	default:
		body, err = sonic.Marshal(in)
	}

	if err != nil {
		return err
	}

	req.Header.SetContentTypeBytes(unsafex.String2Bytes(ct))
	req.SetBodyRaw(body)

	return nil
}

func (c *Client) ResponseDecoder(ctx context.Context, res *protocol.Response, out any) error {
	name := transport.ContentSubtype(unsafex.Bytes2String(res.Header.ContentType()))
	var err error
	switch name {
	case "x-protobuf":
		err = proto.Unmarshal(res.Body(), out.(proto.Message))
	case "plain":
		*(out.(*string)) = string(res.Body())
	default:
		err = sonic.Unmarshal(res.Body(), out)
	}
	if err != nil {
		return err
	}
	return nil
}
