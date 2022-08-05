package gold

import (
	"log"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
	e *Engine
}

func (c *Context) HTML(code int, html string) {
	c.W.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.W.WriteHeader(code)
	_, err := c.W.Write([]byte(html))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Context) Render(code int, name string, data any) {
	c.W.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.W.WriteHeader(code)
	err := c.e.Render(c.W, name, data, c)		
	if err != nil {
		log.Fatal(err)
	}
}
