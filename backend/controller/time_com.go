package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/gin-gonic/gin"
)

// POST /time_coms
func CreateTime_com(c *gin.Context) {
	var time_com entity.Time_com
	if err := c.ShouldBindJSON(&time_com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&time_com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": time_com})
}

// GET /time_com/:id
func GetTime_com(c *gin.Context) {
	var time_com entity.Time_com
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM time_coms WHERE id = ?", id).Find(&time_com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time_com})
}

// GET /times
func ListTime_coms(c *gin.Context) {
	var time_coms []entity.Time_com
	if err := entity.DB().Raw("SELECT * FROM time_coms").Find(&time_coms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time_coms})
}

// DELETE /time_coms/:id
func DeleteTime_com(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM time_coms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time_com not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /time_coms
func UpdateTime_com(c *gin.Context) {
	var time_com entity.Time_com
	if err := c.ShouldBindJSON(&time_com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", time_com.ID).First(&time_com); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	if err := entity.DB().Save(&time_com).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time_com})
}
