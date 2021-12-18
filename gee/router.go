package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, patternn string, handler HandlerFunc) {
	key := method + "-" + patternn
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
		log.Printf("Client %16s Route %4s - %s Query %s",
			ctx.Req.RemoteAddr,
			ctx.Method,
			ctx.Path,
			ctx.Req.URL.RawQuery)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
