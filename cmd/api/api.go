package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhetkerbaevan/green-api/internal/service"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer { //Create new instance of APIServer
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter() //Create a new router

	handler := service.NewHandler()
	handler.RegisterRoutes(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}