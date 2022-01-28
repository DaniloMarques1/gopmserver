package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/danilomarques1/gopmserver/handler"
	"github.com/danilomarques1/gopmserver/repository"
	"github.com/danilomarques1/gopmserver/service"
	"github.com/danilomarques1/gopmserver/util"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

const tables = `
	CREATE TABLE IF NOT EXISTS master(
		id VARCHAR(32) PRIMARY KEY,
		email VARCHAR(100) UNIQUE NOT NULL,
		pwd_hash VARCHAR(100) NOT NULL
	);
`

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
	if _, err := server.db.Exec(tables); err != nil {
		log.Fatal(err)
	}

	server.router.Use(middleware)
	masterRepository := repository.NewMasterRepository(server.db)
	masterService := service.NewMasterService(masterRepository)
	masterHandler := handler.NewMasterHandler(masterService)

	server.router.Post("/master", masterHandler.Save)
	server.router.Post("/session", masterHandler.Session)

	// handler for only authorized routes
	authGroup := server.router.Group(nil)
	authGroup.Use(authMiddleware)

	authGroup.Get("/password", masterHandler.GetPassword)
}

func (server *Server) Start() {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := util.GetTokenFromHeader(r.Header.Get("Authorization"))
		if err != nil {
			util.RespondERR(w, err)
			return
		}
		id, err := util.VerifyToken(token)
		if err != nil {
			util.RespondERR(w, err)
			return
		}
		r.Header.Add("userId", id)
		next.ServeHTTP(w, r)
	})
}
