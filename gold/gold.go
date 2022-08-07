package gold

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Engine struct {
	router
	Renderer
	pool sync.Pool
}

func New() *Engine {
	engine := &Engine{
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
	engine.pool.New = func() any {
		return engine.allocateContext()
	}
	return engine
}

func (e *Engine) allocateContext() any {
	return &Context{e: e}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := e.pool.Get().(*Context)
	c.W = w
	c.R = r
	e.pool.Put(c)
	method := r.Method
	for _, rg := range e.routerGroups {
		node, urlPattern, params := rg.trieNode.Find(r.URL.Path)
		c.pathCache = params
		if node != nil {
			handlerFunc, ok := rg.handlerFuncMap[urlPattern][ANY]
			if ok {
				middlewareFuncs := rg.midderwareFuncMap[urlPattern][ANY]
				rg.Handle(c, handlerFunc, middlewareFuncs...)
				return
			}
			handlerFunc, ok = rg.handlerFuncMap[urlPattern][method]
			if ok {
				middlewareFuncs := rg.midderwareFuncMap[urlPattern][method]
				rg.Handle(c, handlerFunc, middlewareFuncs...)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, r.URL.Path+" "+method+" is not allowed\n")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, r.URL.Path+" is not found\n")
}

func (e *Engine) Run() {
	http.Handle("/", e)
	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
