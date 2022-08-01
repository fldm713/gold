package gold

import (
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

func (rg *routerGroup) Head(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodHead, handlerFunc)
}

func (rg *routerGroup) Post(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodPost, handlerFunc)
}

func (rg *routerGroup) Put(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodPut, handlerFunc)
}

func (rg *routerGroup) Patch(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodPatch, handlerFunc)
}

func (rg *routerGroup) Delete(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodDelete, handlerFunc)
}

func (rg *routerGroup) Trace(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodTrace, handlerFunc)
}

func (rg *routerGroup) Options(routeName string, handlerFunc HandlerFunc) {
	rg.Add(routeName, http.MethodOptions, handlerFunc)
}