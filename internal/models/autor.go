package models

// Autor representa un autor con sus datos y sus libros.
type Autor struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Libros []Libro `json:"libros"`
}
