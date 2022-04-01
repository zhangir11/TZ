package server

import (
	service "authentication/pkg/servicev2"
	storage "authentication/pkg/storagev2"
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer  *http.Server
	authService service.Service
}

//-----------------------------------------------------------------------------------------------

func NewServer() *Server {
	db := storage.SetupConnection()

	sessionManager := storage.NewSessionManager(db, "sessions")
	return &Server{
		authService: service.NewAuthService(sessionManager),
	}
}

//-----------------------------------------------------------------------------------------------

func (s *Server) StartServer(listenAddress string) error {
	router := http.NewServeMux()

	log.Printf("Server launched: %s ... ", listenAddress)

	RegisterHandlers(router, s.authService)

	s.httpServer = &http.Server{
		Addr:    listenAddress,
		Handler: router,
	}
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal("Unable to launch a code: " + err.Error())
	}

	return s.httpServer.Shutdown(ctx)
}
