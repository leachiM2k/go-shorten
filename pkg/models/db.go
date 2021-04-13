package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *DB

func InitDB(driver, dsn string, log *log.Logger) {
	var err error
	db, err = NewDB(driver, dsn, log)
	if err != nil {
		log.Panic(err)
	}

	if err = db.db.Ping(); err != nil {
		log.Panic(err)
	}
}

type DB struct {
	db  *sql.DB
	dsn string
	log *log.Logger
}

func NewDB(driver, dsn string, log *log.Logger) (*DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return &DB{
		db:  db,
		dsn: dsn,
		log: log,
	}, nil
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	d.log.Println(query, args)
	return d.db.Exec(query, args...)
}

func (d *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d.log.Println(query, args)
	return d.db.Query(query, args...)
}

func (d *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	d.log.Println(query, args)
	return d.db.QueryRow(query, args...)
}
