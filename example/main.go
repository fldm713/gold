package main

import (
	"fmt"
	"net/http"

	"github.com/fldm713/gold"
)

func specific(next gold.HandlerFunc) gold.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pre specific middleware\n")
		next(w, r)
		fmt.Fprintf(w, "post specific middleware\n")
	}
}

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
	userGroup.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User: Welcome to the golden era!\n")
	}, specific)
	engine.Run()

}
