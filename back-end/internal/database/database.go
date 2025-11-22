package database

import (
	"log/slog"
	"sync"
)

type DB struct {
	mu    sync.Mutex
	Table map[any][]any
}

func NewDB(table map[any][]any) (*DB, error) {
	return &DB{Table: table}, nil
}

func (db *DB) Close() error {
	slog.Info("Closing database connection")
	return nil
}

func (db *DB) Health() error {
	return nil
}

func (db *DB) BeginTx() error {
	db.mu.Lock()
	return nil
}

func (db *DB) Commit() error {
	db.mu.Unlock()
	return nil
}
