package main

import (
	"fmt"
	"net/http"

	"github.com/fldm713/gold"
)

func main() {
	engine := gold.New()
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Root: Welcome to the golden era!\n")
	})
	engine.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get: Welcome to the golden era!\n")
	})
	engine.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Post: Welcome to the golden era!\n")
	})
	userGroup := engine.Group("user")
	userGroup.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, welcome to the golden era!\n")
	})
	userGroup.Post("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, here is the info you need!\n")
	})
	engine.Any("/any", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ANY: Welcome to the golden nation!\n")
	})
	engine.Post("/any", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Post: Welcome to the golden nation!\n")
	})
	engine.Run()
}
