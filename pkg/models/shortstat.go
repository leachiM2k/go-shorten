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

func GetShortStats(shortenerId int) ([]*ShortStatItem, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, shortenerid, clientip, useragent, referer, createdat "+
			"FROM "+shortstatTableName+" "+
			"WHERE id = $1",
		shortenerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]*ShortStatItem, 0)
	for rows.Next() {
		item := new(ShortStatItem)
		err := rows.Scan(
			&item.ID,
			&item.ShortenerID,
			&item.ClientIP,
			&item.UserAgent,
			&item.Referer,
			&item.CreatedAt)
		if err != nil {
			return nil, err
		}
		stats = append(stats, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}
