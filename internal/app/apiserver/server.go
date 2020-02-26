package apiserver

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	router            *mux.Router
	store             store.Store
	log               *logrus.Logger
	assetPath         string
	previewsDirectory string
	filesDirecotry    string
}

func newServer(store store.Store, log *logrus.Logger, previewsDirectory string, filesDirecotry string) *server {
	s := &server{
		router:            mux.NewRouter(),
		log:               log,
		store:             store,
		previewsDirectory: previewsDirectory,
		filesDirecotry:    filesDirecotry,
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

	s.router.HandleFunc("/v1.0/ping/", s.handlePing()).Methods("GET")
	s.router.HandleFunc("/v1.0/types/", s.handleTypesGet()).Methods("GET")
	s.router.HandleFunc("/v1.0/types/", s.handleRemoveAllTypes()).Methods("DELETE")
	s.router.HandleFunc("/v1.0/types/create/", s.handleCreateType()).Methods("POST")
	s.router.HandleFunc("/v1.0/types/{id}/books/", s.handleGetBooksIDs()).Methods("GET")

	s.router.HandleFunc("/v1.0/books/remove/", s.handleRemoveAllBooks()).Methods("DELETE")

	s.router.HandleFunc("/v1.0/books/create/", s.handleCreateBook()).Methods("POST")
	s.router.HandleFunc("/v1.0/books/{id}/", s.handleGetBookByID()).Methods("GET")

	s.router.HandleFunc("/v1.0/books/{id}/preview/", s.handleBookPreviewUpload()).Methods("POST")
	s.router.HandleFunc("/v1.0/books/{id}/preview/", s.handleBookPreview()).Methods("GET")

	s.router.HandleFunc("/v1.0/books/{id}/file/", s.handleBookFileUpload()).Methods("POST")
	s.router.HandleFunc("/v1.0/books/{id}/file/", s.handleBookFile()).Methods("GET")

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

func (s *server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleRemoveAllBooks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.store.Books().RemoveAll(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleCreateBook() http.HandlerFunc {

	type Response struct {
		BookID      int    `json:"book_id"`
		FileUUID    string `json:"file_uuid"`
		PreviewUUID string `json:"preview_uuid"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		book := &model.Book{}

		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
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

func (s *server) handleGetBookByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])

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

		book.Sanitaize()
		s.respond(w, r, http.StatusOK, book)
	}
}

func (s *server) handleBookPreviewUpload() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Header.Get("X-Token-UUID")

		token, err := s.store.TokensPreview().GetByUID(uuid)

		if token == nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("preview")

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		fileName := handler.Filename

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		savePath := s.assetPath + string(os.PathSeparator) + s.previewsDirectory + string(os.PathSeparator) + fileName
		err = ioutil.WriteFile(savePath, fileBytes, 0644)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		if err := s.store.Books().UpdatePreviewPath(token.BookID, fileName); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err != s.store.TokensPreview().Remove(token) {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}

func (s *server) handleBookPreview() http.HandlerFunc {

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

		previewPath := book.PreviewPath
		file, err := os.Open(previewPath)

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respondFile(w, r, "preview", file)
	}
}
func (s *server) handleBookFile() http.HandlerFunc {

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

		filePath := book.FilePath
		file, err := os.Open(filePath)

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respondFile(w, r, filePath, file)
	}
}

func (s *server) handleBookFileUpload() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Header.Get("X-Token-UUID")

		token, err := s.store.TokensPreview().GetByUID(uuid)

		if token == nil {
			s.error(w, r, http.StatusBadRequest, nil)
			return
		}

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("file")

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		fileName := handler.Filename

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		savePath := s.assetPath + string(os.PathSeparator) + s.filesDirecotry + string(os.PathSeparator) + fileName
		err = ioutil.WriteFile(savePath, fileBytes, 0644)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		if err := s.store.Books().UpdateFilePath(token.BookID, fileName); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err != s.store.TokensPreview().Remove(token) {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}

func (s *server) handleCreateType() http.HandlerFunc {

	type Response struct {
		ID int `json:"ID"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		_type := &model.Type{}

		if err := json.NewDecoder(r.Body).Decode(_type); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		if err := s.store.Types().Create(_type); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		res := &Response{ID: _type.ID}
		s.respond(w, r, http.StatusCreated, res)
	}
}

func (s *server) handleRemoveAllTypes() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := s.store.Types().RemoveAll()

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		s.respond(w, r, http.StatusOK, nil)
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

func (s *server) handleGetBooksIDs() http.HandlerFunc {

	type Response struct {
		BooksIDs []int `json:"books_ids"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		_type, err := s.store.Types().GetByID(id)

		if _type == nil {
			s.error(w, r, http.StatusNotFound, nil)
		}

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, nil)
		}

		books, err := s.store.Books().GetByType(_type)
		s.log.Println(books)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		response := Response{}
		for _, book := range books {
			response.BooksIDs = append(response.BooksIDs, book.ID)
		}

		s.respond(w, r, http.StatusOK, response)
	}
}

func (s *server) respondFile(w http.ResponseWriter, r *http.Request, name string, data io.ReadSeeker) {

	modtime := time.Now()
	w.Header().Add("Content-Disposition", "Attachment")

	http.ServeContent(w, r, name, modtime, data)
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
