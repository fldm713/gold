package gold

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func (c *Context) JSON(code int, data any) {
	c.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(code)
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.W.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Context) XML(code int, data any) {
	c.W.Header().Set("Content-Type", "application/xml; charset=UTF-8")
	c.W.WriteHeader(code)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.W.Write(xmlData)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Context) File(code int, fileName string) {
	http.ServeFile(c.W, c.R, fileName)
}

func (c *Context) Attachment(code int, fileName string, name string) {
	if isASCII(name) {
		c.W.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		c.W.Header().Set("Content-Disposition", `attachment; filename*=UTf-8''`+url.QueryEscape(name))
	}
	http.ServeFile(c.W, c.R, fileName)
}

func (c *Context) Redirect(code int, url string) {
	http.Redirect(c.W, c.R, url, code)
}

func (c *Context) String(code int, format string, values ...any) {
	c.W.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	c.W.WriteHeader(code)
	if len(values) > 0 {
		_, err := fmt.Fprintf(c.W, format, values...)
		if err != nil {
			log.Fatal(err)
		} else {
			c.W.Write([]byte(format))
		}
	} 
}
