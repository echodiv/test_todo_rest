package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	db            *sql.DB
	dbUrl         string
	tagRepository *TagsRepository
}

// Create new storage
func New(db_address string) *Store {
	return &Store{dbUrl: db_address}
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
