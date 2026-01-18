package category

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	repo GormRepository
}

func NewCategoryHandler(repo GormRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req Category
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

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	category, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, category)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
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

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func (h *CategoryHandler) GetAllCategories(context *gin.Context) {
	categories, err := h.repo.GetAll()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, categories)
}
