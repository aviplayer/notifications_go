package main

import (
	. "andrew.com/notifications/cmd/app/accessors"
	"andrew.com/notifications/cmd/app/api"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

const port = 8989

var (
	corsAllowHeaders = "Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Max-Age, Content-Type"
	corsAllowMethods = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin  = "*"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		next(ctx)
	}
}

func main() {
	dbConn := CreateConnection()
	emailAccessor := EmailAccessor{Conn: &dbConn}
	r := api.CreateRouter(&emailAccessor)
	s := fasthttp.Server{
		Handler: CORS(r.GetHandler()),
	}

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	log.Printf("Server starting on port %d\n", port)
	err := s.ListenAndServe(addr)
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
