package gold

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{
			routerGroups: []*routerGroup{
				{
					name:           "",
					handlerFuncMap: make(map[string]map[string]HandlerFunc),
					trieNode:       &trieNode{name: "", children: make([]*trieNode, 0), pattern: ""},
				},
			},
		},
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, rg := range e.routerGroups {
		node, uriPattern := rg.trieNode.Find(r.RequestURI)
		if node != nil {
			handlerFunc, ok := rg.handlerFuncMap[uriPattern][ANY]
			if ok {
				rg.Handle(w, r, handlerFunc)
				return
			}
			handlerFunc, ok = rg.handlerFuncMap[uriPattern][method]
			if ok {
				rg.Handle(w, r, handlerFunc)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, r.RequestURI+" "+method+" is not allowed\n")
			return	
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, r.RequestURI+" is not found\n")
}

func (e *Engine) Run() {
	http.Handle("/", e)
	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
