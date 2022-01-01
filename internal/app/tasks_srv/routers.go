package tasks_srv

import (
	"io"
	"net/http"
)

// .
func (s *Server) getAllTags() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getAllTags start")
		defer func() { s.logger.Debug("--> inside getAllTags end") }()

		io.WriteString(rw, "Get all tags")
	}
}

// .
func (s *Server) getAllTasks() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getAllTasks start")
		defer func() { s.logger.Debug("--> inside getAllTasks end") }()

		io.WriteString(rw, "pass")
	}
}

// .
func (s *Server) createTag() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside createTag start")
		defer func() { s.logger.Debug("--> inside createTag end") }()

		io.WriteString(rw, "Get all tags")
	}
}

// .
func (s *Server) createTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside createTask start")
		defer func() { s.logger.Debug("--> inside createTask end") }()

		io.WriteString(rw, "pass")
	}
}

// .
func (s *Server) getTaskById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getTaskById start")
		defer func() { s.logger.Debug("--> inside getTaskById end") }()

		io.WriteString(rw, "pass")
	}
}

// .
func (s *Server) getTasksByTag() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getTasksByTag start")
		defer func() { s.logger.Debug("--> inside getTasksByTag end") }()

		io.WriteString(rw, "pass")
	}
}
