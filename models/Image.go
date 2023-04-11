package models

import "time"

type Image struct {
	Id          uint64    `json:"Id"`
	Name        string    `json: "Name"`
	ContentType string    `json:"ContentType"`
	Url         string    `json: "Url"`
	Size        int64     `json:"Size"`
	ProductID   uint64    `json:"ProductID,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
