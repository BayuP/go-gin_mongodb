package models

import (
	"time"
)

//Base ...
type Base struct {
	CreatedTime time.Time `json:"created_time"`
	CreatedBy   string    `json:"created_by"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
	DeletedTime time.Time `json:"deleted_time,omitempty"`
	DeletedBy   string    `json:"deleted_by,omitempty"`
}
