package metrics

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"

	"gitlab.com/gobl/gobl/pkg/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

// Server is a prometheus metrics server.
type Server struct {
	mu      sync.Mutex
	cfg     Config
	srv     *http.Server
	reg     *prometheus.Registry
	log     *zap.Logger
	started bool
}

var (
	ErrMetricsDisabled   = errors.New("metrics are disabled")
	ErrMetricsRunning    = errors.New("metrics server is already running")
	ErrMetricsNotRunning = errors.New("metrics server is not running")
)

// DefaultServer creates a new prometheus metrics server with the given config.
// It will create a new HTTP server, handler and prometheus registry, with default handler options.
func DefaultServer(cfg Config) *Server {
	reg := prometheus.NewRegistry()

	httpSrv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		ReadTimeout:       cfg.HTTPServerTimeout,
		ReadHeaderTimeout: cfg.HTTPServerReadHeaderTimeout,
	}

	return WithRegistry(cfg, httpSrv, reg, promhttp.HandlerOpts{})
}

// WithRegistry creates a new prometheus metrics server with the given config, http server,
// prometheus registry and handler options for more control.
func WithRegistry(cfg Config, srv *http.Server, reg *prometheus.Registry, opts promhttp.HandlerOpts) *Server {
	handler := promhttp.HandlerFor(reg, opts)

	mux := http.NewServeMux()
	mux.Handle(cfg.Path, handler)

	srv.Handler = mux

	return &Server{
		cfg:     cfg,
		srv:     srv,
		reg:     reg,
		log:     logger.Logger().With(zap.String("service", "metrics")),
		started: false,
	}
}

// Start starts the metrics server. If the server is already running,
// or is disabled in the application configurations, it will return an error.
func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.cfg.Enabled {
		return ErrMetricsDisabled
	}

	if s.started {
		return ErrMetricsRunning
	}

	go func() {
		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.log.Fatal("Metrics server stopped", zap.Error(err))
		}
	}()

	s.started = true

	return nil
}

// Stop stops the metrics server. If the server is not running,
// or is disabled in the application configurations, it will return an error.
func (s *Server) Stop() error {
	if !s.cfg.Enabled {
		return ErrMetricsDisabled
	}

	if !s.started {
		return ErrMetricsNotRunning
	}

	return s.srv.Close()
}

// Registry returns the prometheus registry used by the metrics server.
func (s *Server) Registry() *prometheus.Registry {
	return s.reg
}

// HTTPServer returns the http server used by the metrics server.
func (s *Server) HTTPServer() *http.Server {
	return s.srv
}

// Unregister unregisters the given prometheus collector from the metrics server.
func (s *Server) Unregister(c prometheus.Collector) bool {
	return s.reg.Unregister(c)
}

// Started returns true if the metrics server is running.
func (s *Server) Started() bool {
	return s.started
}

// Register registers the given instrumentation with the metrics server.
func (s *Server) Register(instrumentation *Instrumentation) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range instrumentation.Collectors() {
		if err := s.reg.Register(c); err != nil {
			return fmt.Errorf("failed to register collector: %w", err)
		}
	}

	return nil
}
