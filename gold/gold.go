package gold

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name             string
	handlerFuncMap   map[string]HandlerFunc
}

func (rg *routerGroup) Add(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
}

type router struct {
	routerGroups []*routerGroup
}

func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:           name,
		handlerFuncMap: make(map[string]HandlerFunc),
	}
	r.routerGroups = append(r.routerGroups, routerGroup)
	return routerGroup
}

func (r *router) Add(name string, handlerFunc HandlerFunc) {
	r.routerGroups[0].Add(name, handlerFunc)
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
					handlerFuncMap: make(map[string]HandlerFunc),
				},
			},
		},
	}
}

func (e *Engine) Run() {
	for _, rg := range e.routerGroups {
		for k, v := range rg.handlerFuncMap {
			if rg.name == "" {
				http.HandleFunc(k, v)
			} else {
				http.HandleFunc("/"+rg.name+k, v)
			}
		}
	}
	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
