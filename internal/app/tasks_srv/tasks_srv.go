package tasks_srv

import (
	"net/http"

	"github.com/echodiv/test_todo_rest/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	CONF_ADDRESS    string = "0.0.0.0:5000"
	CONF_DB_ADDRESS string = "host=localhost port=5432 dbname=dev sslmode=disable user=admin password=secret12345"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store.New(CONF_DB_ADDRESS),
	}
}

func (s *Server) StartServer() error {
	s.logger.Info("Starting API server")
	s.logger.SetLevel(logrus.DebugLevel)
	s.configureRouter()
	if err := s.store.Open(); err != nil {
		return err
	}
	return http.ListenAndServe(CONF_ADDRESS, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/tags", s.getAllTags()).Methods("GET")
	s.router.HandleFunc("/tasks", s.getAllTasks()).Methods("GET")
	s.router.HandleFunc("/tags", s.createTag()).Methods("POST")
	s.router.HandleFunc("/tasks", s.createTask()).Methods("POST")
	s.router.HandleFunc("/tasks/{tagId}", s.getTaskById()).Methods("GET")
	s.router.HandleFunc("/tags/{tagId}", s.getTasksByTag()).Methods("GET")
}
