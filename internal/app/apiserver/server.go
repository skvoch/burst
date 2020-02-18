package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/model"
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
	log.Println("CONFIGURE")
	s.router.HandleFunc("/types/", s.handleTypesGet()).Methods("GET")
	s.router.HandleFunc("/books/", s.handleBooksGet()).Methods("GET")
}

func (s *server) handleTypesGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		types, err := s.store.Types().GetAll()

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, types)
	}
}

func (s *server) handleBooksGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_type := &model.Type{}

		if err := json.NewDecoder(r.Body).Decode(_type); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		books, err := s.store.Books().GetByType(_type)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, books)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
