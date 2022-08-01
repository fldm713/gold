package gold

import "log"

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

func (r *router) Head(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Head(name, handlerFunc)
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

func (r *router) Put(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Put(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Patch(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Patch(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Delete(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Delete(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Trace(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Trace(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}

func (r *router) Options(name string, handlerFunc HandlerFunc) {
	if len(r.routerGroups) > 0 && r.routerGroups[0].name == "" {
		r.routerGroups[0].Options(name, handlerFunc)
	} else {
		log.Fatal("Root group not initialized")
	}
}