package server

import (
	"net/http"

	"github.com/Mungrel/over-calc/auth"

	"github.com/Mungrel/over-calc/handle"

	"github.com/Mungrel/over-calc/db"

	"github.com/bouk/httprouter"
)

// New returns a new server capable of serving the API.
func New() http.Handler {
	authenticatedRouter := httprouter.New()

	authenticatedRouter.POST("/api/operand", handle.CreateOperand)
	authenticatedRouter.GET("/api/operands", handle.ListOperands)
	authenticatedRouter.GET("/api/history", handle.ListHistory)
	authenticatedRouter.POST("/api/eval", handle.Eval)

	authenticator := auth.NewAuthenticator(authenticatedRouter)

	openRouter := httprouter.New()

	openRouter.POST("/api/sign_up", handle.SignUp)
	openRouter.POST("/api/sign_in", handle.SignIn)

	openRouter.NotFound = authenticator

	return &Server{
		handler: openRouter,
	}
}

// Server is a type capable of serving the API.
type Server struct {
	handler http.Handler
}

const acceptedHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

// ServeHTTP implements the http.Handler interface.
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = db.ContextWithDB(ctx, db.Client(ctx))

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", acceptedHeaders)
	server.handler.ServeHTTP(w, r.WithContext(ctx))
}
