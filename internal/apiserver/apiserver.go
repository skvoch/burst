package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Start ...
func (s *APIServer) Start() error {
	err := s.configureLogger()
	s.configureLogger()
	if err != nil {
		return err
	}

	s.logger.Info("Starting API server..")

	return http.ListenAndServe(s.config.BindAdd, s.router)
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return nil
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() error {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *APIServer) handleHello() http.HandleFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
