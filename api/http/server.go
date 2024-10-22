package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	foldercontroller "rollout/internal/controller/folder"
)

type Server struct {
	Router *mux.Router
}

func NewServer(folderCtrl *foldercontroller.Controller) *Server {
	router := mux.NewRouter()
	registerRoutes(router, folderCtrl)
	return &Server{Router: router}
}

func (s *Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on port %s", address)
	return http.ListenAndServe(address, s.Router)
}

func registerRoutes(r *mux.Router, folderCtrl *foldercontroller.Controller) {
	r.HandleFunc("/folders", folderCtrl.CreateFolder).Methods("POST")
}
