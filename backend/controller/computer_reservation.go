package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Computer_reservation
func CreateComputer_reservation(c *gin.Context) {

	var user entity.User
	var computer_reservation entity.Computer_reservation
	var computer entity.Computer
	// var computer_os entity.Computer_os
	var time_com entity.Time_com

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร computer_reservation
	if err := c.ShouldBindJSON(&computer_reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ยูสเซอร์ not found"})
		return
	}

	// 10: ค้นหา COMPUTER ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.Computer_id).First(&computer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	// 11: ค้นหา Time ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.Time_com_id).First(&time_com); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	// 12: สร้าง Computer_reservation
	cr := entity.Computer_reservation{
		User:     user,                      // โยงความสัมพันธ์กับ Entity User
		Computer: computer,                  // โยงความสัมพันธ์กับ Entity Computer
		Time_com: time_com,                  // โยงความสัมพันธ์กับ Entity Time_com
		Date:     computer_reservation.Date, // ตั้งค่าฟิลด์ watchedTime ยัง

	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(cr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&cr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cr})
}

// GET /computer_reservation/:id
func GetComputer_reservation(c *gin.Context) {
	var computer_reservation entity.Computer_reservation
	id := c.Param("id")
	if err := entity.DB().Preload("User").Preload("Computer").Preload("Time_com").Raw("SELECT * FROM Computer_reservation WHERE id = ?", id).Find(&computer_reservation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": computer_reservation})
}

// GET /computer_reservations
func ListComputer_reservations(c *gin.Context) {
	var computer_reservations []entity.Computer_reservation
	if err := entity.DB().Preload("User").Preload("Computer").Preload("Time_com").Raw("SELECT * FROM computer_reservations").Find(&computer_reservations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_reservations})
}

// DELETE /computer_reservations/:id
func DeleteComputer_reservation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM computer_reservations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer_reservation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /computer_reservations
func UpdateComputer_reservation(c *gin.Context) {
	var computer_reservation entity.Computer_reservation
	if err := c.ShouldBindJSON(&computer_reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", computer_reservation.ID).First(&computer_reservation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer_reservation not found"})
		return
	}

	if err := entity.DB().Save(&computer_reservation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_reservation})
}
