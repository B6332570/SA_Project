package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/gin-gonic/gin"
)

// POST /computers
func CreateCOMPUTER(c *gin.Context) {
	var computer entity.COMPUTER
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": computer})
}

// GET /computers/:id
func GetCOMPUTER(c *gin.Context) {
	var computer entity.COMPUTER
	id := c.Param("id")
	if err := entity.DB().Preload("COMPUTER_OS").Raw("SELECT * FROM computers WHERE id = ?", id).Find(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer})
}

// GET /computers
func ListCOMPUTERs(c *gin.Context) {
	var computers []entity.COMPUTER
	if err := entity.DB().Preload("COMPUTER_OS").Raw("SELECT * FROM computers").Find(&computers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computers})
}

// DELETE /computers/:id
func DeleteCOMPUTER(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM computers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /computers
func UpdateCOMPUTER(c *gin.Context) {
	var computer entity.COMPUTER
	if err := c.ShouldBindJSON(&computer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", computer.ID).First(&computer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	if err := entity.DB().Save(&computer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer})
}
