package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jiayishen21/resume-comp-backend/service/education"
	"github.com/jiayishen21/resume-comp-backend/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)

	educationStore := education.NewStore(s.db)
	educationHandler := education.NewHandler(educationStore)

	userHandler.RegisterRoutes(subrouter)
	educationHandler.RegisterRoutes(subrouter)

	log.Println("Started server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
