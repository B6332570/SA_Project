package controller

import (
	"net/http"

	"github.com/B6332570/SA_Project/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Computer_Reservation
func CreateCOMPUTER_RESERVATION(c *gin.Context) {

	var user entity.User
	var computer_reservation entity.COMPUTER_RESERVATION
	var computer entity.COMPUTER
	var computer_os entity.COMPUTER_OS
	var time entity.TIME

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร computer_reservation
	if err := c.ShouldBindJSON(&computer_reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 10: ค้นหา COMPUTER_OS ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.COMP_OS_ID).First(&computer_os); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer_os not found"})
		return
	}

	// 11: ค้นหา COMPUTER ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.COMP_ID).First(&computer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer not found"})
		return
	}

	// 12: ค้นหา Time ด้วย id
	if tx := entity.DB().Where("id = ?", computer_reservation.TIME_ID).First(&time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	// 13: สร้าง COMPUTER_RESERVATION
	cr := entity.COMPUTER_RESERVATION{
		User: user,                      // โยงความสัมพันธ์กับ Entity User
		C:    computer,                  // โยงความสัมพันธ์กับ Entity COMPUTER
		CO:   computer_os,               // ตั้งค่าฟิลด์ COMPUTER_OS
		TIME: time,                      // โยงความสัมพันธ์กับ Entity Time
		Date: computer_reservation.Date, // ตั้งค่าฟิลด์ watchedTime

	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(cr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 14: บันทึก
	if err := entity.DB().Create(&cr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cr})
}

// GET /computer_reservation/:id
func GetCOMPUTER_RESERVATION(c *gin.Context) {
	var computer_reservation entity.COMPUTER_RESERVATION
	id := c.Param("id")
	if err := entity.DB().Preload("User").Preload("COMPUTER_OS").Preload("COMPUTER").Preload("TIME").Raw("SELECT * FROM watch_videos WHERE id = ?", id).Find(&computer_reservation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": computer_reservation})
}

// GET /computer_reservations
func ListCOMPUTER_RESERVATIONs(c *gin.Context) {
	var computer_reservations []entity.COMPUTER_RESERVATION
	if err := entity.DB().Preload("User").Preload("COMPUTER_OS").Preload("COMPUTER").Preload("TIME").Raw("SELECT * FROM computer_reservations").Find(&computer_reservations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": computer_reservations})
}

// DELETE /computer_reservations/:id
func DeleteCOMPUTER_RESERVATION(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM computer_reservations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "computer_reservation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /computer_reservations
func UpdateCOMPUTER_RESERVATION(c *gin.Context) {
	var computer_reservation entity.COMPUTER_RESERVATION
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
