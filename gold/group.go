package gold

import (
	"log"
	"net/http"
)

const ANY string = "ANY"

type HandlerFunc func(c *Context)

type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc

type routerGroup struct {
	name              string
	handlerFuncMap    map[string]map[string]HandlerFunc
	trieNode          *trieNode
	middlewares       []MiddlewareFunc
	midderwareFuncMap map[string]map[string][]MiddlewareFunc
}

func (rg *routerGroup) Use(middleWareFuncs ...MiddlewareFunc) {
	rg.middlewares = append(rg.middlewares, middleWareFuncs...)
}

func (rg *routerGroup) Handle(c *Context, h HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	for _, m := range rg.middlewares {
		h = m(h)
	}
	for _, m := range middlewareFuncs {
		h = m(h)
	}
	h(c)
}

func (rg *routerGroup) Add(routeName string, method string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	var urlPattern string
	if rg.name == "" {
		urlPattern = routeName
	} else {
		urlPattern = "/" + rg.name + routeName
	}
	if _, ok := rg.handlerFuncMap[urlPattern]; !ok {
		rg.handlerFuncMap[urlPattern] = make(map[string]HandlerFunc)
		rg.midderwareFuncMap[urlPattern] = make(map[string][]MiddlewareFunc)
	}
	if _, ok := rg.handlerFuncMap[urlPattern][method]; ok {
		log.Fatalf("Repeated, url %s, method: %s\n", urlPattern, method)
	}
	rg.handlerFuncMap[urlPattern][method] = handlerFunc
	rg.midderwareFuncMap[urlPattern][method] = append(rg.midderwareFuncMap[urlPattern][method], middlewareFuncs...)
	rg.trieNode.Insert(urlPattern)
}

func (rg *routerGroup) Any(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, ANY, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Get(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodGet, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Head(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodHead, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Post(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodPost, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Put(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodPut, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Patch(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodPatch, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Delete(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodDelete, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Trace(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodTrace, handlerFunc, middlewareFuncs...)
}

func (rg *routerGroup) Options(routeName string, handlerFunc HandlerFunc, middlewareFuncs ...MiddlewareFunc) {
	rg.Add(routeName, http.MethodOptions, handlerFunc, middlewareFuncs...)
}
