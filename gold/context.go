package gold

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)


const defaultMaxMemory int64 = 32 << 20 // 32 MB

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	e          *Engine
	queryCache url.Values
	formCache  url.Values
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
		}
	} else {
		_, err := c.W.Write([]byte(format))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Context) initQueryCache() {
	if c.R != nil {
		c.queryCache = c.R.URL.Query()
	} else {
		c.queryCache = url.Values{}
	}
}

func (c *Context) QueryParam(name string) string {
	c.initQueryCache()
	return c.queryCache.Get(name)
}

func (c *Context) QueryParams() url.Values {
	c.initQueryCache()
	return c.queryCache
}

func (c *Context) initPostFormCache() {
	if c.R != nil {
		if err := c.R.ParseMultipartForm(defaultMaxMemory); err != nil {
			if !errors.Is(err, http.ErrNotMultipart) {
				log.Fatal(err) 
			}
		}
		c.formCache = c.R.PostForm
	} else {
		c.formCache = url.Values{}
	}
}

func (c *Context) FormValue(name string) string {
	c.initPostFormCache()
	return c.formCache.Get(name)
}

func (c *Context) FormValues() url.Values {
	c.initPostFormCache()
	return c.formCache
}

func (c *Context) FormFile(name string) *multipart.FileHeader {
		file, header, err := c.R.FormFile(name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		return header
}
