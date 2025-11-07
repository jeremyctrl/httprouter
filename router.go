package httprouter

import "net/http"

type Handler func(w http.ResponseWriter, req *http.Request, params ...string)

type Router struct {
	routes []routeDef
	groups mphGroups
}

func New() *Router {
	return &Router{
		routes: make([]routeDef, 0),
		groups: nil,
	}
}

func (r *Router) GET(path string, handler Handler) *Router {
	return r.Handle(http.MethodGet, path, handler)
}

func (r *Router) HEAD(path string, handler Handler) *Router {
	return r.Handle(http.MethodHead, path, handler)
}

func (r *Router) POST(path string, handler Handler) *Router {
	return r.Handle(http.MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler Handler) *Router {
	return r.Handle(http.MethodPut, path, handler)
}

func (r *Router) DELETE(path string, handler Handler) *Router {
	return r.Handle(http.MethodDelete, path, handler)
}

func (r *Router) OPTIONS(path string, handler Handler) *Router {
	return r.Handle(http.MethodOptions, path, handler)
}

func (r *Router) PATCH(path string, handler Handler) *Router {
	return r.Handle(http.MethodPatch, path, handler)
}

func (r *Router) Handle(method, path string, handler Handler) *Router {
	r.routes = append(r.routes, routeDef{
		method,
		path,
		handler,
	})
	return r
}

func (r *Router) Build() *Router {
	r.groups = build(r.routes)
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
