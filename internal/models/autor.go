package models

type Autor struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string  `gorm:"not null" json:"name"`
	Email  string  `gorm:"not null;unique" json:"email"`
	Libros []Libro `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"libros"`
}
