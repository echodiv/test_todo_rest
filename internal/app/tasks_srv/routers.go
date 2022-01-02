package tasks_srv

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all tags from storage
func (s *Server) getAllTags() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getAllTags start")
		defer func() { s.logger.Debug("--> inside getAllTags end") }()
		listOfTags, err := s.store.Tag().GetAllTags()
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			return
		}
		response, err := json.Marshal(listOfTags)
		if err != nil {
			rw.WriteHeader((http.StatusBadGateway))
			return
		}
		io.WriteString(rw, string(response))
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

// Get all tasks by tag id
func (s *Server) getTasksByTag() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Debug("--> inside getTasksByTag start")
		defer func() { s.logger.Debug("--> inside getTasksByTag end") }()
		vars := mux.Vars(r)
		tagId, err := strconv.Atoi(vars["tagId"])
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			io.WriteString(rw, err.Error())
			return
		}
		tasksByTag, err := s.store.Tag().FindById(tagId)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			io.WriteString(rw, err.Error())
			return
		}
		respose, err := json.Marshal(tasksByTag)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			io.WriteString(rw, err.Error())
			return
		}
		io.WriteString(rw, string(respose))
	}
}
