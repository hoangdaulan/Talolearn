package course

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service CourseService
}

func NewCourseHandler(service CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req Course
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateCourse(&req); err != nil {
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
	course, err := h.service.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(200, course)
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.service.ListCourses()
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

	if err := h.service.UpdateCourse(uint(id), updates); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, _ := h.service.GetCourseByID(uint(id))
	c.JSON(200, res)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteCourse(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Course deleted successfully"})
}

func (h *CourseHandler) ListCourses(c *gin.Context) {
	courses, err := h.service.ListCourses()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, courses)
}

func (h *CourseHandler) AddTeacherToCourse(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course id"})
		return
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid teacher id"})
		return
	}

	if err := h.service.AddTeacherToCourse(uint(courseID), uint(teacherID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Teacher added into course successfully"})
}

func (h *CourseHandler) RemoveTeacherFromCourse(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course id"})
		return
	}

	teacherID, err := strconv.ParseUint(c.Param("teacher_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid teacher id"})
		return
	}

	if err := h.service.RemoveTeacherFromCourse(uint(courseID), uint(teacherID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Teacher removed from course successfully"})
}
