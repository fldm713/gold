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

func Hello(c *gold.Context) {
	c.Render(http.StatusOK, "hello", "World")
}

type User struct {
	Name string `json:"name" xml:"name"`
}

func main() {
	engine := gold.New()
	// user := &struct {
	// 	Name string `xml:"name"`
	// } {
	// 	Name: "User1",
	// }
	user := &User{
		Name: "User1",
	}
	engine.Get("/json", func(c *gold.Context) {
		c.JSON(http.StatusOK, user)
	})
	engine.Get("/redirect", func(c *gold.Context) {
		c.Redirect(http.StatusFound, "/json")
	})
	engine.Run()

}
