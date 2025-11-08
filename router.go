// Package httprouter provides a fast, minimal-perfect-hash-based HTTP router.
//
// An example is:
//
// package main
//
// import (
//
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/jeremyctrl/httprouter"
//
// )
//
//	func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//		fmt.Fprint(w, "Welcome!\n")
//	}
//
//	func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//	}
//
//	func main() {
//		router := httprouter.New().
//			GET("/", Index).
//			GET("/hello/:name", Hello).
//			Build()
//
//		log.Fatal(http.ListenAndServe(":8080", router))
//	}
//
// The router dispatches requests based on both HTTP method and URL path.
// When a route matching the method and path is found, the corresponding
// handler is invoked to serve the request.
package httprouter

import "net/http"

type Handler func(w http.ResponseWriter, req *http.Request, params Params)

type Router struct {
	routes   []routeDef
	groups   mphGroups
	NotFound http.HandlerFunc
}

func New() *Router {
	return &Router{
		routes: make([]routeDef, 0),
		groups: nil,
	}
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) GET(path string, handler Handler) *Router {
	return r.Handle(http.MethodGet, path, handler)
}

// HEAD is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) HEAD(path string, handler Handler) *Router {
	return r.Handle(http.MethodHead, path, handler)
}

// POST is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) POST(path string, handler Handler) *Router {
	return r.Handle(http.MethodPost, path, handler)
}

// PUT is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) PUT(path string, handler Handler) *Router {
	return r.Handle(http.MethodPut, path, handler)
}

// DELETE is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) DELETE(path string, handler Handler) *Router {
	return r.Handle(http.MethodDelete, path, handler)
}

// OPTIONS is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) OPTIONS(path string, handler Handler) *Router {
	return r.Handle(http.MethodOptions, path, handler)
}

// PATCH is a shortcut for router.Handle(http.MethodGet, path, handler)
func (r *Router) PATCH(path string, handler Handler) *Router {
	return r.Handle(http.MethodPatch, path, handler)
}

// Handle registers a new request handler for the given path and method.
func (r *Router) Handle(method, path string, handler Handler) *Router {
	r.routes = append(r.routes, routeDef{
		method,
		path,
		handler,
	})
	return r
}

// Build compiles all registered routes into minimal-perfect-hash tables.
// Must be called after route registration and before ServeHTTP.
func (r *Router) Build() *Router {
	r.groups = build(r.routes)
	return r
}

// ServeHTTP dispatches an incoming request to the matching route handler.
// If no route matches, the NotFound handler is called (or http.NotFound by default).
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route, params := find(r.groups, req.Method, req.URL.Path)
	if route == nil {
		if r.NotFound != nil {
			r.NotFound(w, req)
		} else {
			http.NotFound(w, req)
		}
		return
	}
	route.handler(w, req, params)
}
