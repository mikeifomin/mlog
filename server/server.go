package server

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

type Server struct {
	Bind   string
	Access map[string]string
	Tokens []string

	PgHost string
	PgUser string
	PgPass string
	PgName string
	PgPort uint16

	AdminDir string

	adminPrefix  []byte
	adminHandler fasthttp.RequestHandler
}

func (s *Server) Run() {
	s.adminPrefix = []byte("/admin/")
	s.adminHandler = fasthttp.FSHandler(s.AdminDir, 1)

	err := fasthttp.ListenAndServe(s.Bind, s.mainHandler)
	if err != nil {
		panic(err)
	}
}

func (s *Server) mainHandler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	switch {
	//case "/":
	//common.PlaintextHandler(ctx)
	//case "/query":
	//common.JSONHandler(ctx)
	case bytes.HasPrefix(path, s.adminPrefix):
		s.adminHandler(ctx)
	default:
		ctx.Error("unexpected path", fasthttp.StatusBadRequest)
	}
}
