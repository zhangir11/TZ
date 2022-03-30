package server

import (
	"authentication/pkg/service"
	"authentication/pkg/storage"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer  *http.Server
	authService service.Service
}

func NewServer() *Server {
	db := storage.SetupConnection()

	sessionManager := storage.NewSessionManager(db, "sessions")
	return &Server{
		authService: service.NewAuthService(sessionManager),
	}
}

func (s *Server) StartServer(listenAddress string) error {
	router := mux.NewRouter()

	log.Printf("Server launched: %s ... ", listenAddress)

	RegisterHandlers(router, s.authService)

	s.httpServer = &http.Server{
		Addr:    listenAddress,
		Handler: router,
	}

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal("Unable to launch a code: " + err.Error())
	}

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}
