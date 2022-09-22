package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	httpServer *http.Server
}

func New(listen string) (*Server, error) {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "pong", http.StatusOK)
	})
	h := h2c.NewHandler(mux, &http2.Server{})

	s.httpServer = &http.Server{
		Addr:              listen,
		Handler:           h,
		ReadHeaderTimeout: time.Second,
	}

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	l, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return err
	}
	log.Printf("Listening on: %v", l.Addr())

	idleConnsClosed := make(chan struct{})
	go func() {
		<-ctx.Done()
		ctxT, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		s.httpServer.Shutdown(ctxT)

		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	err = s.httpServer.Serve(l)
	<-idleConnsClosed
	return err
}
