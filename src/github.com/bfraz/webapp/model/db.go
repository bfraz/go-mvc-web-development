package model

import "database/sql"

var db DB

type DB interface {
  Query(string, ...interface{}) (Rows, error)
}

type Rows interface {
  Next() bool
  Scan(dest ...interface{}) error
  Close() error
}

type sqlDB struct {
	db *sql.DB
}
func (s sqlDB) Query(query string, args ...interface{}) (Rows, error ){
  return s.db.Query(query, args...)
}

func SetDatabase(database *sql.DB) {
	db = &sqlDB{database}
}
