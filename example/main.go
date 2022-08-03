package main

import (
	"fmt"
	"net/http"

	"github.com/fldm713/gold"
)

func main() {
	engine := gold.New()
	engine.Use(func(next gold.HandlerFunc) gold.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Pre middleware\n")
			next(w, r)
			fmt.Fprintf(w, "Post middleware\n")
		}
	})
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Root: Welcome to the golden era!\n")
	})
	engine.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Match all: Welcome to the golden era!\n")
	})
	engine.Get("/:param", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param: Welcome to the golden era!\n")
	})
	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Static: Welcome to the golden era!\n")
	})
	userGroup := engine.Group("users")
	userGroup.Get("/:*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Match all: Hi user, welcome to the golden era!\n")
	})
	userGroup.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param: Hi user, welcome to the golden era!\n")
	})
	userGroup.Get("/:id/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, here is the info you need!\n")
	})

	engine.Run()

}
