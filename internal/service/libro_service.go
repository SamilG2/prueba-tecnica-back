package service

import (
	"errors"
	"prueba-tecnica-back/internal/models"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateLibro(libro *models.Libro) (*models.Libro, error) {
	if libro.Title == "" {
		return nil, errors.New("El libro debe tener un título")
	}
	if libro.AuthorID == 0 {
		return nil, errors.New("El libro debe estar asociado a un autor válido")
	}

	err := s.db.Create(libro).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Preload("Autor").First(libro, libro.ID).Error

	return libro, err
}

func (s *Service) GetAllLibros() ([]*models.Libro, error) {
	var libros []*models.Libro
	err := s.db.Preload("Autor").Find(&libros).Error
	return libros, err
}

func (s *Service) GetLibroByID(id uint) (*models.Libro, error) {
	var libro models.Libro
	err := s.db.Preload("Autor").First(&libro, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("libro no encontrado")
		}
		return nil, err
	}
	return &libro, nil
}

func (s *Service) UpdateLibro(id uint, nuevosDatos *models.Libro) (*models.Libro, error) {
	var libro models.Libro

	if err := s.db.First(&libro, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("libro no encontrado")
		}
		return nil, err
	}

	err := s.db.Model(&libro).Updates(nuevosDatos).Error
	if err != nil {
		return nil, err
	}

	s.db.Preload("Autor").First(&libro, id)

	return &libro, nil
}

func (s *Service) DeleteLibro(id uint) error {
	var libro models.Libro
	err := s.db.First(&libro, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("libro no encontrado")
		}
		return err
	}
	return s.db.Delete(&libro).Error
}
