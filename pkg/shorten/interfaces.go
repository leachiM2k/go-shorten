package shorten

import "time"

type Backend interface {
	Create(request CreateRequest) (*Entity, error)
	Read(code string) (*Entity, error)
	Update(entity *Entity) (*Entity, error)
	Delete(code string) error
}

type Entity struct {
	Owner      string                 `json:"owner"`
	Link       string                 `json:"link"`
	Code       string                 `json:"code"`
	Count      int                    `json:"count"`
	MaxCount   int                    `json:"maxCount"`
	CreatedAt  time.Time              `json:"createdAt"`
	UpdatedAt  time.Time              `json:"updatedAt"`
	StartTime  *time.Time             `json:"startTime"`
	ExpiresAt  *time.Time             `json:"expiresAt"`
	Attributes *map[string]interface{} `json:"attributes"`
}

type CreateRequest struct {
	Owner      *string                `json:"owner"`
	Link       *string                `json:"link"`
	Code       string                 `json:"code,omitempty"`
	MaxCount   int                    `json:"maxCount,omitempty"`
	StartTime  *time.Time             `json:"startTime,omitempty"`
	ExpiresAt  *time.Time             `json:"expiresAt,omitempty"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

type UpdateRequest struct {
	Link       string                 `json:"link"`
	MaxCount   int                    `json:"maxCount,omitempty"`
	StartTime  *time.Time             `json:"startTime,omitempty"`
	ExpiresAt  *time.Time             `json:"expiresAt,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}
