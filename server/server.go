package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/danilomarques1/gopmserver/handler"
	"github.com/danilomarques1/gopmserver/repository"
	"github.com/danilomarques1/gopmserver/service"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	router *chi.Mux
	db     *sql.DB
}

// TODO here we are using an http connection
// we may in the future use a raw socket connection
func NewServer() (*Server, error) {
	db, err := sql.Open("sqlite3", "gopm.db")
	if err != nil {
		return nil, err
	}
	server := Server{db: db}
	server.router = chi.NewRouter()
	return &server, nil
}

func (server *Server) Init() {
	server.router.Use(middleware)
	masterRepository := repository.NewMasterRepository(server.db)
	masterService := service.NewMasterService(masterRepository)
	_ = handler.NewMasterHandler(masterService)
}

func (server *Server) Start() {
	port := "8080"
	log.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server.router)) // TODO change port
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
