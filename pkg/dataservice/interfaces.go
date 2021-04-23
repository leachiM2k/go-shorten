package dataservice

import "time"

type Entity struct {
	ID          int                     `json:"-"`
	Owner       string                  `json:"owner"`
	Link        string                  `json:"link"`
	Code        string                  `json:"code"`
	Description string                  `json:"description"`
	Count       int                     `json:"count"`
	MaxCount    int                     `json:"maxCount"`
	CreatedAt   time.Time               `json:"createdAt"`
	UpdatedAt   time.Time               `json:"updatedAt"`
	StartTime   *time.Time              `json:"startTime"`
	ExpiresAt   *time.Time              `json:"expiresAt"`
	Attributes  *map[string]interface{} `json:"attributes"`
}

type StatEntity struct {
	ShortenerID int       `json:"shortenerID"`
	ClientIP    string    `json:"clientIP"`
	UserAgent   string    `json:"userAgent"`
	Referer     string    `json:"referer"`
	CreatedAt   time.Time `json:"timestamp"`
}

type CreateRequest struct {
	Owner       *string                 `json:"owner"`
	Link        *string                 `json:"link"`
	Code        string                  `json:"code,omitempty"`
	Description string                  `json:"description,omitempty"`
	MaxCount    int                     `json:"maxCount,omitempty"`
	StartTime   *time.Time              `json:"startTime,omitempty"`
	ExpiresAt   *time.Time              `json:"expiresAt,omitempty"`
	Attributes  *map[string]interface{} `json:"attributes,omitempty"`
}

type UpdateRequest struct {
	Owner       *string                `json:"owner"`
	Link        string                 `json:"link"`
	Description string                 `json:"description,omitempty"`
	MaxCount    int                    `json:"maxCount,omitempty"`
	StartTime   *time.Time             `json:"startTime,omitempty"`
	ExpiresAt   *time.Time             `json:"expiresAt,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}
