package models

import "time"

type Image struct {
	Id          uint64    `json:"Id"`
	Url         string    `json: "Url"`
	ContentType string    `json:"ContentType"`
	Data        string    `json:"Data"`
	Size        int64     `json:"Size"`
	ProductID   uint64    `json:"ProductID,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
