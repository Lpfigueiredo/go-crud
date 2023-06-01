package pkg

import "github.com/google/uuid"

type Book struct {
	Id     uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	ISBN   string    `json:"isbn"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Price  float64   `json:"price" binding:"required,numeric"`
}
