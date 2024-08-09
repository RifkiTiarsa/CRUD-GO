package main

import (
	"challenge-goapinew/config"
	"challenge-goapinew/controller"
	"challenge-goapinew/repository"
	"challenge-goapinew/usecase"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server struct {
	uc     usecase.LoundryUseCase
	router *mux.Router
	host   string
}

func (s *Server) initRoute() {
	s.router = mux.NewRouter()

	customerController := controller.NewLoundryController(s.uc)
	customerController.Route(s.router)
}

func (s *Server) Run() {
	s.initRoute()

	fmt.Printf("Server running on host %s\n", s.host)
	err := http.ListenAndServe(s.host, s.router)
	if err != nil {
		log.Fatalf("Server not running on host %s, because error %v", s.host, err)
	}
}
func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBNAME)
	fmt.Printf("DSN: %s\n", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	fmt.Println("Connection to database successful")

	repository := repository.NewLoundryRepository(db)
	usecase := usecase.NewLoundryUseCase(repository)

	host := cfg.ServerAddress
	return &Server{
		uc:   usecase,
		host: host,
	}
}
