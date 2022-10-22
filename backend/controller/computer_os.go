package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/gin-gonic/gin"
)

// POST /computer_OSs
func CreateComputer_os(c *gin.Context) {
	var computer_os entity.Computer_os
	if err := c.ShouldBindJSON(&computer_os); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&computer_os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": computer_os})
}

// GET /computer_os/:id
func GetComputer_os(c *gin.Context) {
	var computer_os entity.Computer_os
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM computer_os WHERE id = ?", id).Find(&computer_os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_os})
}

// GET /computer_oss
func ListComputer_oss(c *gin.Context) {
	var computer_oss []entity.Computer_os
	if err := entity.DB().Raw("SELECT * FROM computer_os").Find(&computer_oss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_oss})
}

// DELETE /computer_oss/:id
func DeleteComputer_os(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM computer_os WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /computer_oss
func UpdateComputer_os(c *gin.Context) {
	var computer_os entity.Computer_os
	if err := c.ShouldBindJSON(&computer_os); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", computer_os.ID).First(&computer_os); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer_os not found"})
		return
	}

	if err := entity.DB().Save(&computer_os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_os})
}
