package models

import "time"

type Product struct {
	ID          uint64    `json:"id" gorm:"AUTO_INCREMENT"`
	Name        string    `json:"Name"`
	Description string    `json:"Description,omitempty"`
	Images      *[]Image  `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time `json:"created_at"`
}
