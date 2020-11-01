package router

import (
	"github.com/go-chi/chi"
	"github.com/tonquangtu/http_server/internal/handler"
	"github.com/tonquangtu/http_server/internal/repository"
	"net/rpc/jsonrpc"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	registerRoute(r)
	return r
}

func registerRoute(r *chi.Mux) {
	client, _ := jsonrpc.Dial("tcp", "127.0.0.1:6969")
	userRepository := repository.NewUserRepository(client)
	userHandler := handler.NewUserHandler(userRepository)

	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.ServeHTTP)
	})

}
