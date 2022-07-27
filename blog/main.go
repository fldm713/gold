package main

import (
	"fmt"
	"net/http"

	"github.com/fldm713/gold"
)

func main() {
	engine := gold.New()
	engine.Any("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the golden era!\n")
	})
	userGroup := engine.Group("user")
	userGroup.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, welcome to the golden era!\n")
	})
	userGroup.Post("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, here is the info you need!\n")
	})
	engine.Run()
}
