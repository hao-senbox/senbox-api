package common

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 30 * time.Second
	_defaultWriteTimeout    = 30 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 5 * time.Second
)

// HTTP -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// NewLogger -.
func NewServer(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
