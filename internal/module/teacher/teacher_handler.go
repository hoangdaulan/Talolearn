package teacher

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	service TeacherService
}

func NewTeacherHandler(service TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var req TeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	if _, err := h.service.CreateTeacher(uint(id), req.Bio, req.CategoryID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetTeacherByUserID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, res)
}

func (h *TeacherHandler) GetTeacherByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	teacher, err := h.service.GetTeacherByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(200, teacher)
}

func (h *TeacherHandler) GetTeacherByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	teacher, err := h.service.GetTeacherByUserID(uint(userID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(200, teacher)
}

func (h *TeacherHandler) FindTeacherByName(c *gin.Context) {
	name := c.Query("name")
	teacher, err := h.service.FindTeacherByName(name)
	if err != nil {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(200, teacher)
}

func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	var teachers []Teacher
	teachers, err := h.service.GetAllTeachers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, teachers)
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTeacher(uint(id), updates); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, _ := h.service.GetTeacherByID(uint(id))
	c.JSON(200, res)
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteTeacher(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
