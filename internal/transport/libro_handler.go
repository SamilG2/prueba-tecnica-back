package transport

import (
	"net/http"
	"strconv"

	"prueba-tecnica-back/internal/models"
	"prueba-tecnica-back/internal/service"

	"github.com/gin-gonic/gin"
)

type LibroHandler struct {
	service *service.Service
}

func New(s *service.Service) *LibroHandler {
	return &LibroHandler{service: s}
}

func (h *LibroHandler) Create(c *gin.Context) {
	var libro models.Libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	created, err := h.service.CreateLibro(&libro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *LibroHandler) GetAll(c *gin.Context) {
	libros, err := h.service.GetAllLibros()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, libros)
}

func (h *LibroHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	libroID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	libro, err := h.service.GetLibroByID(uint(libroID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, libro)
}

func (h *LibroHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	libroID, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if libroID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser mayor que 0"})
		return
	}

	var libro models.Libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	updated, err := h.service.UpdateLibro(uint(libroID), &libro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *LibroHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	libroID, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if libroID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser mayor que 0"})
		return
	}

	if err := h.service.DeleteLibro(uint(libroID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
