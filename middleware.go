package main

import (
	"context"
	"net/http"
)

type CustomMux struct {
	http.ServerMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleWare(next func(http.Handler)) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServerMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

func MiddleWareUtility(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request))
	ctx := r.Context()

	if ctx == nil {
		ctx = context.Background()
	}
	from := r.Header.Get("Referer")
	if from == "" {
		from = r.Host
	}

}
