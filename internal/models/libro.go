package models

type Libro struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	AuthorID    uint    `gorm:"not null" json:"authorId"`
	Autor       *Autor  `gorm:"foreignKey:AuthorID" json:"autor"`
}
