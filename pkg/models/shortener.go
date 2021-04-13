package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"time"
)

type ShortenerItem struct {
	ID         int
	Owner      string
	Code       string
	Link       string
	Count      int
	MaxCount   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StartTime  *time.Time
	ExpiresAt  *time.Time
	Attributes *Attributes
}

type Attributes map[string]interface{}

// Make the Attributes struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Attributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attributes struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Attributes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func AddShort(item ShortenerItem) error {
	_, err := db.Exec(
		"INSERT INTO shortener (owner, code, link, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		item.Owner, item.Code, item.Link, item.Count, item.MaxCount, item.CreatedAt, item.UpdatedAt, item.StartTime, item.ExpiresAt, item.Attributes)
	if err != nil {
		return err
	}

	return nil
}

func AllShortenerByOwner(owner string) ([]*ShortenerItem, error) {
	rows, err := db.Query("SELECT id, owner, code, link, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes FROM shortener WHERE owner = $1", owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make([]*ShortenerItem, 0)
	for rows.Next() {
		item := new(ShortenerItem)
		err := rows.Scan(&item.ID, &item.Owner, &item.Code, &item.Link, &item.Count, &item.MaxCount, &item.CreatedAt, &item.UpdatedAt, &item.StartTime, &item.ExpiresAt, &item.Attributes)
		if err != nil {
			return nil, err
		}
		teams = append(teams, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}

func GetShortenerByCode(code string) (*ShortenerItem, error) {
	item := new(ShortenerItem)
	err := db.QueryRow("SELECT id, owner, code, link, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes FROM shortener WHERE code = $1", code).Scan(&item.ID, &item.Owner, &item.Code, &item.Link, &item.Count, &item.MaxCount, &item.CreatedAt, &item.UpdatedAt, &item.StartTime, &item.ExpiresAt, &item.Attributes)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func UpdateShortener(item *ShortenerItem) (*ShortenerItem, error) {
	_, err := db.Exec(
		"UPDATE shortener SET "+
			"link = $1, "+
			"maxCount = $2, "+
			"updatedAt = $3, "+
			"startTime = $4, "+
			"expiresAt = $5, "+
			"attributes = $6 "+
			"WHERE code = $7",
		item.Link,
		item.MaxCount,
		time.Now(),
		item.StartTime,
		item.ExpiresAt,
		item.Attributes,
		item.Code)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteShortenerByCode(code string) error {
	_, err := db.Exec("DELETE FROM shortener WHERE code = $1", code)
	if err != nil {
		return err
	}

	return nil
}
