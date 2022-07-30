package gold

import (
	"fmt"
	"log"
	"net/http"
)

const ANY string = "ANY"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name             string
	handlerFuncMap   map[string]map[string]HandlerFunc
}

func (rg *routerGroup) Add(routeName string, method string, handlerFunc HandlerFunc) {
	if _, ok := rg.handlerFuncMap[routeName]; !ok {
		rg.handlerFuncMap[routeName] = make(map[string]HandlerFunc)
	}
	var uri string
	if rg.name == "" {
		uri = routeName
	} else {
		uri = "/" + rg.name + routeName
	}
	if _, ok := rg.handlerFuncMap[routeName][method]; ok {
		log.Fatalf("Repeated, uri: %s, method: %s\n", uri + routeName, ANY)
	}
	rg.handlerFuncMap[routeName][method] = handlerFunc
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
		name:             name,
		handlerFuncMap:   make(map[string]map[string]HandlerFunc),
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
					name:             "",
					handlerFuncMap:   make(map[string]map[string]HandlerFunc),
				},
			},
		},
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, rg := range e.routerGroups {
		for routeName, methodMap := range rg.handlerFuncMap {
			var uri string
			if rg.name == "" {
				uri = routeName
			} else {
				uri = "/" + rg.name + routeName
			}
			if r.RequestURI == uri {
				handlerFunc, ok := methodMap[ANY]
				if ok {
					handlerFunc(w, r)
					return
				}
				handlerFunc, ok = methodMap[method]
				if ok {
					handlerFunc(w, r)
					return
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, uri+" "+method+" is not allowed\n")
				return
			}
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
