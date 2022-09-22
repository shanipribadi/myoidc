package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/zeebo/blake3"
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Listen   string
	Secret   string
	Issuer   string
	Insecure bool
}

func New(cfg *Config) (*Server, error) {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "pong", http.StatusOK)
	})

	storage := newStorage()
	key := blake3.Sum256([]byte(cfg.Secret))
	provider, err := newOP(context.Background(), storage, key, cfg.Insecure, cfg.Issuer)
	if err != nil {
		return nil, err
	}
	mux.Handle("/", provider.HttpHandler())

	h := h2c.NewHandler(mux, &http2.Server{})

	s.httpServer = &http.Server{
		Addr:              cfg.Listen,
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
