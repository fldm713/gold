package main

import (
	"fmt"
	"net/http"

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
	engine.Get("/user/:id", func(c *gold.Context) {
		c.String(http.StatusOK, "id: %s", c.PathParam("id"))
	})
	// engine.Get("/order/:name/info/:id", func(c *gold.Context) {
	// 	c.String(http.StatusOK, "Order name: %s, info id %s", c.PathParam("name"), c.PathParam("id"))
	// })
	engine.Get("/order/:name/info/:id", func(c *gold.Context) {
		c.String(http.StatusOK, "Path parameters %v", c.PathParams())
	})
	engine.Run()

}
