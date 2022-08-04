package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/fldm713/gold"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c *gold.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c *gold.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	engine := gold.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	engine.Renderer = t

	engine.Get("/", func(c *gold.Context) error {
		return c.HTML(http.StatusOK, "<h1>html</h1><br>")
	})

	engine.Get("/hello", Hello)
	
	engine.Run()

}
