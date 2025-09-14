{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{- range .MethodSets}}
const Operation{{$svrType}}{{.OriginalName}} = "/{{$svrName}}/{{.OriginalName}}"
{{- end}}

type {{.ServiceType}}HertzServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

type {{.ServiceType}}HertzMiddleware interface {
{{- range .MethodSets}}
	{{.Name}}Middleware() []app.HandlerFunc
{{- end}}
}

func Register{{.ServiceType}}HertzServer(s *hertz.Server, srv {{.ServiceType}}HertzServer, mw {{.ServiceType}}HertzMiddleware, routeMw... app.HandlerFunc) {
	r := s.Group("/", routeMw...)
	{{- range .Methods}}
	r.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_Hertz_Handler(s, srv, mw)...)
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_Hertz_Handler(s *hertz.Server, srv {{$svrType}}HertzServer, mw {{$svrType}}HertzMiddleware) []app.HandlerFunc {
	var handlers []app.HandlerFunc
	if mw != nil {
		handlers = append(handlers, mw.{{.Name}}Middleware()...)
	}
	h := func(ctx context.Context, c *app.RequestContext) {
		var err error
		var req {{.Request}}
		err = c.BindAndValidate(&req)
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
		resp, err := srv.{{.Name}}(ctx, &req)
		if err != nil {
			s.ErrorEncoder(ctx, c, err)
			return
		}
		s.ResponseEncoder(ctx, c, resp)
		return
	}

	return append(handlers, h)
}
{{end}}

type {{.ServiceType}}HertzClient interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(ctx context.Context, req *{{.Request}}, opts ...config.RequestOption) (rsp *{{.Reply}}, err error)
{{- end}}
}

type {{.ServiceType}}HertzClientImpl struct{
	cc *hertz.Client
	host string
}

func New{{.ServiceType}}HertzClient (client *hertz.Client, host string) {{.ServiceType}}HertzClient {
	return &{{.ServiceType}}HertzClientImpl{client, host}
}

{{range .MethodSets}}
    {{- if ne .Comment ""}}
    {{.Comment}}
    {{- end}}
func (c *{{$svrType}}HertzClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...config.RequestOption) (*{{.Reply}}, error) {
	uri := c.host + "{{.Path}}"
	var out {{.Reply}}

	err := c.cc.Do(ctx, uri, "{{.Method}}", in, &out, opts...)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
{{end}}

type {{.ServiceType}}HertzLocalImpl struct{
	impl {{.ServiceType}}HertzServer
}

func New{{.ServiceType}}HertzLocalImpl (impl {{.ServiceType}}HertzServer) {{.ServiceType}}HertzClient {
	return &{{.ServiceType}}HertzLocalImpl{impl}
}

{{range .MethodSets}}
    {{- if ne .Comment ""}}
    {{.Comment}}
    {{- end}}
func (c *{{$svrType}}HertzLocalImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...config.RequestOption) (*{{.Reply}}, error) {
	return c.impl.{{.Name}}(ctx, in)
}
{{end}}