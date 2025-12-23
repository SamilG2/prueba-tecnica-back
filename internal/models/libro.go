package models

type Libro struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	AuthorID    uint    `json:"authorId"`
	Autor       Autor   `json:"autor"`
}
