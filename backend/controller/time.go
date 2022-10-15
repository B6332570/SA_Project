package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/gin-gonic/gin"
)

// POST /times
func CreateTIME(c *gin.Context) {
	var time entity.TIME
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": time})
}

// GET /time/:id
func GetTIME(c *gin.Context) {
	var time entity.TIME
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM times WHERE id = ?", id).Find(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}

// GET /times
func ListTIMEs(c *gin.Context) {
	var times []entity.TIME
	if err := entity.DB().Raw("SELECT * FROM times").Find(&times).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": times})
}

// DELETE /times/:id
func DeleteTIME(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM times WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /time
func UpdateTIME(c *gin.Context) {
	var time entity.TIME
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", time.ID).First(&time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	if err := entity.DB().Save(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}
