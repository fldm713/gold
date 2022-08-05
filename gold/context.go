package gold

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
	e *Engine
}

func (c *Context) HTML(code int, html string) error {
	c.W.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.W.WriteHeader(code)
	_, err := c.W.Write([]byte(html))
	return err
}

func (c *Context) Render(code int, name string, data any) error {
	c.W.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.W.WriteHeader(code)
	err := c.e.Render(c.W, name, data, c)		
	return err
}

func (c *Context) JSON(code int, data any) error {
	c.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(code)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(jsonData)
	return err
}

func (c *Context) XML(code int, data any) error {
	c.W.Header().Set("Content-Type", "application/xml; charset=UTF-8")
	c.W.WriteHeader(code)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(xmlData)
	return err
}

func (c *Context) File(code int, fileName string) error {
	http.ServeFile(c.W, c.R, fileName)
	return nil
}

func (c *Context) Attachment(code int, fileName string, name string) error {
	if isASCII(name) {
		c.W.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		c.W.Header().Set("Content-Disposition", `attachment; filename*=UTf-8''`+url.QueryEscape(name))
	}
	http.ServeFile(c.W, c.R, fileName)
	return nil
}