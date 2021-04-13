package models

import (
	_ "github.com/lib/pq"
	"time"
)

type LogItem struct {
	ID        int
	PNum      string
	Route     string
	Method    string
	Payload   []byte
	Timestamp time.Time
}

func PersistLog(item LogItem) error {
	_, err := db.Exec("INSERT INTO log (pnum, method, route, timestamp, payload) VALUES($1, $2, $3, $4, $5)", item.PNum, item.Method, item.Route, time.Now(), item.Payload)
	if err != nil {
		return err
	}

	return nil
}
