package models

import (
	_ "github.com/lib/pq"
	"time"
)

type ShortStatItem struct {
	ID          int
	ShortenerID int
	ClientIP    string
	UserAgent   string
	Referer     string
	CreatedAt   time.Time
}

const shortstatTableName = "shortstat"

func AddShortStat(item ShortStatItem) error {
	_, err := db.Exec(
		"INSERT INTO "+shortstatTableName+
			" (shortenerid, clientip, useragent, referer, createdat)"+
			" VALUES($1, $2, $3, $4, $5)",
		item.ShortenerID,
		item.ClientIP,
		item.UserAgent,
		item.Referer,
		time.Now())
	if err != nil {
		return err
	}

	return nil
}

func AllShortStats(code string) ([]*ShortStatItem, error) {
	rows, err := db.Query(
		"SELECT"+
			" stat.clientip, stat.useragent, stat.referer, stat.createdat"+
			" FROM "+shortstatTableName+" stat"+
			" LEFT JOIN "+shortenerTableName+" s ON s.id = stat.shortenerid"+
			" WHERE s.code = $1",
		code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	statItems := make([]*ShortStatItem, 0)
	for rows.Next() {
		item := new(ShortStatItem)
		err := rows.Scan(
			&item.ClientIP,
			&item.UserAgent,
			&item.Referer,
			&item.CreatedAt)
		if err != nil {
			return nil, err
		}
		statItems = append(statItems, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return statItems, nil
}
