package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/store"
)

type server struct {
	router *mux.Router
	store  store.Store
	log    *logrus.Logger
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		log:    logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/types", s.handleTypesGet()).Methods("GET")
}

func (s *server) handleTypesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
