package gold

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	router
	Renderer
}

func New() *Engine {
	return &Engine{
		router: router{
			routerGroups: []*routerGroup{
				{
					name:              "",
					handlerFuncMap:    make(map[string]map[string]HandlerFunc),
					midderwareFuncMap: make(map[string]map[string][]MiddlewareFunc),
					trieNode:          &trieNode{name: "", children: make([]*trieNode, 0), pattern: ""},
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
			c := &Context{
				W: w,
				R: r,
				e: e,
			}
			handlerFunc, ok := rg.handlerFuncMap[uriPattern][ANY]
			if ok {
				middlewareFuncs := rg.midderwareFuncMap[uriPattern][ANY]
				err := rg.Handle(c, handlerFunc, middlewareFuncs...)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			handlerFunc, ok = rg.handlerFuncMap[uriPattern][method]
			if ok {
				middlewareFuncs := rg.midderwareFuncMap[uriPattern][method]
				err := rg.Handle(c, handlerFunc, middlewareFuncs...)
				if err != nil {
					log.Fatal(err)
				}
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
