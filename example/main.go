package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fldm713/gold"
)

func main() {
	engine := gold.New()
	engine.Get("/hello", func(c *gold.Context) {
		id := c.QueryParam("id")
		c.String(http.StatusOK, id)
	})
	engine.Get("/user", func(c *gold.Context) {
		c.String(http.StatusOK, "Query parameters: %#v\n", c.QueryParams())
		fmt.Printf("%v", c.QueryParams())
	})
	engine.Post("/hello", func(c *gold.Context) {
		id := c.FormValue("id")
		c.String(http.StatusOK, id)
	})
	engine.Post("/user", func(c *gold.Context) {
		c.String(http.StatusOK, "Form values %v", c.FormValues())
	})
	engine.Post("/file", func(c *gold.Context) {
		file := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()
		dst, err := os.Create("./uploaded/" + file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()
		io.Copy(dst, src)
	})
	engine.Run()

}
