package gold

import (
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2021-01-15/web"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type routerGroup struct {
	name             string
	handlerFuncMap   map[string]HandlerFunc
	handlerMethodMap map[string][]string
}

func (rg *routerGroup) Add(name string, handlerFunc HandlerFunc) {
	rg.handlerFuncMap[name] = handlerFunc
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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.routerGroups {
		for name, methodH
	}
}

func (e *Engine) Run() {
	http.Handle("/", e)
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
