package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	SELECT_TAGS_BY_ID  = "SELECT id, tag_name, created FROM tags where id=$1"
	SELECT_ALL_TAGS    = "SELECT id, tag_name, created FROM tags"
	SELECT_TASK_BY_TAG = "SELECT id, task_name, task_description, created from tasks where tag_id=$1"

	INSERT_NEW_TAG = "INSERT INTO tags (name) VALUES ($1) RETURNING id"
)

type Store struct {
	db             *sql.DB
	dbUrl          string
	logger         *logrus.Logger
	tagRepository  *TagsRepository
	taskRepository *TaskRepository
}

// Create new storage
func New(db_address string, logger *logrus.Logger) *Store {
	return &Store{dbUrl: db_address, logger: logger}
}

// Open connection and ping database
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.dbUrl)
	if err != nil {
		log.Println("error with postgres connection")
		return err
	}
	if err := db.Ping(); err != nil {
		log.Println("Error with postgres ping")
		return err
	}
	s.db = db
	return nil
}

// Close connection to database
func (s *Store) Close() {
	s.db.Close()
}

// Get tagRepository from storage
func (s *Store) Tag() *TagsRepository {
	if s.tagRepository == nil {
		s.tagRepository = &TagsRepository{s}
	}
	return s.tagRepository
}

// Get taskRepository from storage
func (s *Store) Task() *TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = &TaskRepository{s}
	}
	return s.taskRepository
}
