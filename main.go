package main

import (
	"context"
	"io"
	"fmt"
	"errors"
	"time"
	"net/http"
)

type M map [string] interface{}

func main ()
mux := new(CustomMux)
mux.RegisterMiddleWare(MiddleWareUtility)

mux.HandlerFunc("/api/search",func (w http.ResponseWriter, r* http.Request)  {
	from := r.Conten() . value ("from").(string)
	w.Writte([]byte(from))
})

server := new (http.Server)
server.Handler = mux
server.Addr = "80"

fmt.Println("Starting server at ", server . Addr)
server.ListenAndServer()