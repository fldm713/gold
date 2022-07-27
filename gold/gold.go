package gold

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name             string
	handlerFuncMap   map[string]HandlerFunc
	handlerMethodMap map[string][]string
}

func (rg *routerGroup) Any(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
	rg.handlerMethodMap["ANY"] = append(rg.handlerMethodMap["ANY"], name)
}

func (rg *routerGroup) Get(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
	rg.handlerMethodMap[http.MethodGet] = append(rg.handlerMethodMap[http.MethodGet], name)
}

func (rg *routerGroup) Post(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
	rg.handlerMethodMap[http.MethodPost] = append(rg.handlerMethodMap[http.MethodPost], name)
}

type router struct {
	routerGroups []*routerGroup
}

func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:             name,
		handlerFuncMap:   make(map[string]HandlerFunc),
		handlerMethodMap: make(map[string][]string),
	}
	r.routerGroups = append(r.routerGroups, routerGroup)
	return routerGroup
}

func (r *router) Any(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Get(name, handlerFunc)
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
		r.routerGroups[0].Get(name, handlerFunc)
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
					handlerFuncMap:   make(map[string]HandlerFunc),
					handlerMethodMap: make(map[string][]string),
				},
			},
		},
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, rg := range e.routerGroups {
		for route, handlerFunc := range rg.handlerFuncMap {
			var uri string
			if rg.name == "" {
				uri = route
			} else {
				uri = "/" + rg.name + route
			}
			if r.RequestURI == uri {
				names, ok := rg.handlerMethodMap["ANY"]
				if ok {
					for _, name := range names {
						if name == route {
							handlerFunc(w, r)
							return
						}
					}
				}
				names, ok = rg.handlerMethodMap[method]
				if ok {
					for _, name := range names {
						if name == route {
							handlerFunc(w, r)
							return
						}
					}
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
