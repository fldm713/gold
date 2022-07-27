package main

import (
	"fmt"
	"net/http"

	"github.com/fldm713/gold"
)

func main() {
	engine := gold.New()
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the golden era!\n")
	})
	userGroup := engine.Group("user")
	userGroup.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi user, welcome to the golden era!\n")
	})
	engine.Run()
}
