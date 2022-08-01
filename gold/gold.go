package gold

import (
	"fmt"
	"log"
	"net/http"
)

const ANY string = "ANY"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name           string
	handlerFuncMap map[string]map[string]HandlerFunc
	trieNode       *trieNode
}

func (rg *routerGroup) Add(routeName string, method string, handlerFunc HandlerFunc) {
	var uriPattern string
	if rg.name == "" {
		uriPattern = routeName
	} else {
		uriPattern = "/" + rg.name + routeName
	}
	if _, ok := rg.handlerFuncMap[uriPattern]; !ok {
		rg.handlerFuncMap[uriPattern] = make(map[string]HandlerFunc)
	}
	if _, ok := rg.handlerFuncMap[uriPattern][method]; ok {
		log.Fatalf("Repeated, uri %s, method: %s\n", uriPattern, method)
	}
	rg.handlerFuncMap[uriPattern][method] = handlerFunc
	rg.trieNode.Insert(uriPattern)
}

func (rg *routerGroup) Any(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, ANY, handlerFunc)
}

func (rg *routerGroup) Get(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodGet, handlerFunc)
}

func (rg *routerGroup) Post(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodPost, handlerFunc)
}

type router struct {
	routerGroups []*routerGroup
}

func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:           name,
		handlerFuncMap: make(map[string]map[string]HandlerFunc),
		trieNode:       &trieNode{name: "", children: make([]*trieNode, 0), pattern: ""},
	}
	r.routerGroups = append(r.routerGroups, routerGroup)
	return routerGroup
}

func (r *router) Any(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Any(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Get(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Get(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Post(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Post(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

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
				handlerFunc(w, r)
				return
			}
			handlerFunc, ok = rg.handlerFuncMap[uriPattern][method]
			if ok {
				handlerFunc(w, r)
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
