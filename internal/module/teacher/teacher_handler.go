package teacher

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	repo GormRepository
}

func NewTeacherHandler(repo GormRepository) *TeacherHandler {
	return &TeacherHandler{repo: repo}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var req Teacher
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, req)
}

func (h *TeacherHandler) GetTeacherByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	teacher, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(200, teacher)
}

func (h *TeacherHandler) FindTeacherByName(c *gin.Context) {
	name := c.Query("name")
	teacher, err := h.repo.FindByName(name)
	if err != nil {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(200, teacher)
}

func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	var teachers []Teacher
	teachers, err := h.repo.GetAll()
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

	if err := h.repo.UpdatePartial(uint(id), updates); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, _ := h.repo.GetByID(uint(id))
	c.JSON(200, res)
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Deleted": "Successfully"})
}
