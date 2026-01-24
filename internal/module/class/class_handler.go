package class

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	service ClassService
}

func NewClassHandler(service ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TeacherID   uint   `json:"teacher_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	class := Class{
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   req.TeacherID,
	}

	if err := h.service.CreateClass(class); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create class"})
		return
	}

	c.JSON(201, gin.H{"message": "Class created successfully"})
}

func (h *ClassHandler) GetClass(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}

	class, err := h.service.GetClass(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Class not found"})
		return
	}

	c.JSON(200, class)
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TeacherID   uint   `json:"teacher_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}

	if err := h.service.DeleteClass(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete class"})
		return
	}

	c.JSON(200, gin.H{"message": "Class deleted successfully"})
}

func (h *ClassHandler) ListClasses(c *gin.Context) {
	classes, err := h.service.ListClasses()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve classes"})
		return
	}

	c.JSON(200, classes)
}
