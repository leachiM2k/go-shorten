package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	e "github.com/pkg/errors"
	"time"
)

type ShortenerItem struct {
	ID          int
	Owner       string
	Code        string
	Link        string
	Description string
	Count       int
	MaxCount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartTime   *time.Time
	ExpiresAt   *time.Time
	Attributes  *Attributes
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

const shortenerTableName = "shortener"

func AddShort(item ShortenerItem) (*ShortenerItem, error) {
	sqlItem, err := db.Exec(
		"INSERT INTO "+shortenerTableName+
			" (owner, code, link, description, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes)"+
			" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		item.Owner,
		item.Code,
		item.Link,
		item.Description,
		item.Count,
		item.MaxCount,
		item.CreatedAt,
		item.UpdatedAt,
		item.StartTime,
		item.ExpiresAt,
		item.Attributes)
	if err != nil {
		return nil, err
	}

	lastInsertId, _ := sqlItem.LastInsertId()
	item.ID = int(lastInsertId)

	return &item, nil
}

func AllShortsByOwner(owner string) ([]*ShortenerItem, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, owner, code, link, description, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes "+
			"FROM "+shortenerTableName+" "+
			"WHERE owner = $1",
		owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shortenerItems := make([]*ShortenerItem, 0)
	for rows.Next() {
		item := new(ShortenerItem)
		err := rows.Scan(
			&item.ID,
			&item.Owner,
			&item.Code,
			&item.Link,
			&item.Description,
			&item.Count,
			&item.MaxCount,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.StartTime,
			&item.ExpiresAt,
			&item.Attributes)
		if err != nil {
			return nil, err
		}
		shortenerItems = append(shortenerItems, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return shortenerItems, nil
}

func GetShortByCode(code string) (*ShortenerItem, error) {
	item := new(ShortenerItem)
	err := db.QueryRow(
		"SELECT "+
			"id, owner, code, link, description, count, maxCount, createdAt, updatedAt, startTime, expiresAt, attributes "+
			"FROM "+shortenerTableName+" "+
			"WHERE code = $1",
		code).Scan(
		&item.ID,
		&item.Owner,
		&item.Code,
		&item.Link,
		&item.Description,
		&item.Count,
		&item.MaxCount,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.StartTime,
		&item.ExpiresAt,
		&item.Attributes)
	if err != nil {
		if e.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func UpdateShort(item *ShortenerItem) (*ShortenerItem, error) {
	_, err := db.Exec(
		"UPDATE "+shortenerTableName+" SET "+
			"link = $1, "+
			"description = $2, "+
			"maxCount = $3, "+
			"updatedAt = $4, "+
			"startTime = $5, "+
			"expiresAt = $6, "+
			"attributes = $7 "+
			"WHERE code = $8 AND owner = $9",
		item.Link,
		item.Description,
		item.MaxCount,
		time.Now(),
		item.StartTime,
		item.ExpiresAt,
		item.Attributes,
		item.Code,
		item.Owner)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteShortByCode(owner string, code string) error {
	_, err := db.Exec("DELETE FROM "+shortenerTableName+" WHERE owner = $1 AND code = $2", owner, code)
	if err != nil {
		return err
	}

	return nil
}
