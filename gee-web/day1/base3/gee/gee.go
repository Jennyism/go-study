package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, req *http.Request)

type Engine interface {
	GET(pattern string, handler HandleFunc)
	POST(pattern string, handler HandleFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Run(addr string) error
}

type GeeEngine struct {
	router map[string]HandleFunc
}

func New() *GeeEngine {
	return &GeeEngine{make(map[string]HandleFunc)}
}

func (g *GeeEngine) addRouter(method string, pattern string, handler HandleFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	g.router[key] = handler
}

func (g *GeeEngine) GET(pattern string, handler HandleFunc) {
	g.addRouter("GET", pattern, handler)
}

func (g *GeeEngine) POST(pattern string, handler HandleFunc) {
	g.addRouter("POST", pattern, handler)
}

func (g *GeeEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s-%s", req.Method, req.URL.Path)
	if handler, ok := g.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (g *GeeEngine) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}
