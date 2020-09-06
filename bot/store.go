package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	DatabaseURL string
}

type Store struct {
	config *DbConfig
	db     *sql.DB
}

func NewDbConfig(databaseUrl string) *DbConfig {
	return &DbConfig{DatabaseURL: databaseUrl}
}

func NewStore(config *DbConfig) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
