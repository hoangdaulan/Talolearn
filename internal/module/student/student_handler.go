package student

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req Student
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	student, err := h.service.CreateStudent(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(201, student)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid student ID"})
		return
	}
	student, err := h.service.GetStudentByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(200, student)
}
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := h.service.UpdateStudent(idParam, updates)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{"message": "Student updated successfully"})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	err := h.service.DeleteStudent(idParam)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(200, gin.H{"message": "Student deleted successfully"})
}

func (h *StudentHandler) ListStudents(c *gin.Context) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list students"})
		return
	}

	c.JSON(200, students)
}
