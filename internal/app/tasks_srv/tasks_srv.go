package tasks_srv

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *Server) StartServer() error {
	s.logger.Info("Starting API server")
	s.logger.SetLevel(logrus.DebugLevel)
	s.configureRouter()
	return http.ListenAndServe("0.0.0.0:5000", s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/tags", s.getAllTags()).Methods("GET")
	s.router.HandleFunc("/tasks", s.getAllTasks()).Methods("GET")
	s.router.HandleFunc("/tags", s.createTag()).Methods("POST")
	s.router.HandleFunc("/tasks", s.createTask()).Methods("POST")
	s.router.HandleFunc("/tasks/{tagId}", s.getTaskById()).Methods("GET")
	s.router.HandleFunc("/tags/{tagId}", s.getTasksByTag()).Methods("GET")
}
