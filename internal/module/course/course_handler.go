package course

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	repo GormRepository
}

func NewCourseHandler(repo GormRepository) *CourseHandler {
	return &CourseHandler{repo: repo}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req Course
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateCourse(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, req)
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	course, err := h.repo.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(200, course)
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.repo.ListCourses()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, courses)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
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

	if err := h.repo.UpdateCourse(uint(id), updates); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, _ := h.repo.GetCourseByID(uint(id))
	c.JSON(200, res)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteCourse(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Course deleted successfully"})
}

func (h *CourseHandler) ListCourses(c *gin.Context) {
	courses, err := h.repo.ListCourses()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, courses)
}
