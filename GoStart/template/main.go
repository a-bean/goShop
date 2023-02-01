package main

import (
	"bytes"
	"html/template"
	"strings"
)

var tpl = `
type {{.Name}}HttpServer struct {
	server {{$.Name}}Server

	router gin.IRouter
}

func Register{{.Name}}HttpServer(server {{.Name}}Server, router gin.IRouter) {
	//我现在想用gin.Default,如果开发中我想使用qit
	g := &{{.Name}}HttpServer{server: server, router: router}
	g.RegisterService()
}


{{ range .Methods }}
func (g *{{ $.Name }}HttpServer) {{ .HandlerName }}(c *gin.Context) {
	var in {{ .Request }}
	if err := c.BindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := g.server.{{ .Name }}(c, &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
{{ end }}

func (g *{{$.Name}}HttpServer) RegisterService() {
{{ range .Methods }}
	g.router.Handle("{{ .Method }}", "{{ .Path }}", g.{{ .HandlerName }})
{{ end }}
}
`

type serviceDesc struct {
	Name    string
	Methods []method
}

type method struct {
	Name    string
	Request string
	Reply   string

	//http rule
	Path   string
	Method string
	Body   string
}

func (m *method) HandlerName() string {
	return m.Name + "_0"
}

func main() {
	//模板
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}

	s := serviceDesc{
		Name: "Greeter",

		Methods: []method{
			{
				Name:    "SayHello",
				Request: "HelloRequest",
				Reply:   "HelloReply",
				Path:    "/v1/sayhello",
				Method:  "POST",
				Body:    "*",
			},
		},
	}

	err = tmpl.Execute(buf, s)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}
