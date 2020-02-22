package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store"
)

type ctxKey int8

const (
	ctxKeyRequestID ctxKey = iota
)

type server struct {
	router *mux.Router
	store  store.Store
	log    *logrus.Logger
}

func newServer(store store.Store, log *logrus.Logger) *server {
	s := &server{
		router: mux.NewRouter(),
		log:    log,
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/v1.0/books/create", s.handleCreateBook()).Methods("POST")
	s.router.HandleFunc("/v1.0/books/{id}", s.handleCreateBook()).Methods("POST")

	s.log.Info("Endpoints:")
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		s.log.Info(tpl, err1, met, err2)
		return nil
	})
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-request-ID", id)

		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.log.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})

		logger.Infof("started %s, %s", r.Method, r.RequestURI)

		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Infof("finished time %v ", time.Now().Sub(start))

	})
}

func (s *server) handleCreateBook() http.HandlerFunc {

	type Response struct {
		FileUUID    string `file_uuid:"id"`
		PreviewUUID string `preview_uuid:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		book := &model.Book{}

		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		book.Sanitaize()

		if err := s.store.Books().Create(book); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		fileToken := &model.PDFToken{UID: uuid.New().String(), BookID: book.ID}
		previewToken := &model.PreviewToken{UID: uuid.New().String(), BookID: book.ID}

		if err := s.store.TokensPDF().Create(fileToken); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		if err := s.store.TokensPreview().Create(previewToken); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		response := &Response{
			FileUUID:    fileToken.UID,
			PreviewUUID: previewToken.UID,
		}

		s.respond(w, r, http.StatusCreated, response)
	}
}

func (s *server) handleGetBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["ID"])

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		book, err := s.store.Books().GetByID(ID)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)

		}

		if book == nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respond(w, r, http.StatusOK, book)
	}
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

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
