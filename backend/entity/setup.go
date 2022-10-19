package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-05.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&EMPLOYEE{},
		&ROLE{},
		&PROVINCE{},
		&MemberClass{},
		&User{},
		&COMPUTER_OS{},
		&COMPUTER_RESERVATION{},
		&COMPUTER{},
		&TIME{},
	)

	db = database

	//User
	db.Model(&User{}).Create(&User{
		PIN:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		CIV:       "1111111111111",
		PHONE:     "0811111111",
		EMAIL:     "preechapat@mail.com",
		PASSWORD:  "1111111111111",
		ADDRESS:   "ถนน a อำเภอ v",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	db.Model(&User{}).Create(&User{
		PIN:       "B6222222",
		FirstName: "kawin",
		LastName:  "l.pat",
		CIV:       "2222222222222",
		PHONE:     "0922222222",
		EMAIL:     "kawin@mail.com",
		PASSWORD:  "2222222222222",
		ADDRESS:   "หอ b อำเภอ r",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	//COMPUTER_OS data
	comp_os_name1 := COMPUTER_OS{
		Name: "Windows",
	}
	db.Model(&COMPUTER_OS{}).Create(&comp_os_name1)

	comp_os_name2 := COMPUTER_OS{
		Name: "macOS",
	}
	db.Model(&COMPUTER_OS{}).Create(&comp_os_name2)

	//COMPUTER data
	db.Model(&COMPUTER{}).Create(&COMPUTER{
		COMP_NAME:   "W01A",
		COMP_ROOM:   "ROOM A",
		COMPUTER_OS: comp_os_name1,
	})

	db.Model(&COMPUTER{}).Create(&COMPUTER{
		COMP_NAME:   "W02A",
		COMP_ROOM:   "ROOM A",
		COMPUTER_OS: comp_os_name1,
	})

	db.Model(&COMPUTER{}).Create(&COMPUTER{
		COMP_NAME:   "W03A",
		COMP_ROOM:   "ROOM A",
		COMPUTER_OS: comp_os_name1,
	})

	db.Model(&COMPUTER{}).Create(&COMPUTER{
		COMP_NAME:   "M01B",
		COMP_ROOM:   "ROOM B",
		COMPUTER_OS: comp_os_name2,
	})

	db.Model(&COMPUTER{}).Create(&COMPUTER{
		COMP_NAME:   "M02B",
		COMP_ROOM:   "ROOM B",
		COMPUTER_OS: comp_os_name2,
	})

	//TIME data
	time_p1 := TIME{
		TIME_PERIOD: "08:00 - 09:00",
	}
	db.Model(&TIME{}).Create(&time_p1)

	time_p2 := TIME{
		TIME_PERIOD: "09:00 - 10:00",
	}
	db.Model(&TIME{}).Create(&time_p2)

	time_p3 := TIME{
		TIME_PERIOD: "10:00 - 11:00",
	}
	db.Model(&TIME{}).Create(&time_p3)

	time_p4 := TIME{
		TIME_PERIOD: "11:00 - 12:00",
	}
	db.Model(&TIME{}).Create(&time_p4)

	time_p5 := TIME{
		TIME_PERIOD: "12:00 - 13:00",
	}
	db.Model(&TIME{}).Create(&time_p5)

	time_p6 := TIME{
		TIME_PERIOD: "13:00 - 14:00",
	}
	db.Model(&TIME{}).Create(&time_p6)

	time_p7 := TIME{
		TIME_PERIOD: "14:00 - 15:00",
	}
	db.Model(&TIME{}).Create(&time_p7)

	time_p8 := TIME{
		TIME_PERIOD: "15:00 - 16:00",
	}
	db.Model(&TIME{}).Create(&time_p8)

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ COMPUTER มาเก็บไว้ในตัวแปรก่อน
	// cn = comp_name ที่มาจาก COMP_NAME ใน Entity COMPUTER
	var cn1 COMPUTER
	db.Raw("SELECT * FROM computers WHERE name = ?", "W01A").Scan(&cn1)
	var cn2 COMPUTER
	db.Raw("SELECT * FROM computers WHERE name = ?", "W02A").Scan(&cn2)
	var cn3 COMPUTER
	db.Raw("SELECT * FROM computers WHERE name = ?", "W03A").Scan(&cn3)
	var cn4 COMPUTER
	db.Raw("SELECT * FROM computers WHERE name = ?", "M01B").Scan(&cn4)
	var cn5 COMPUTER
	db.Raw("SELECT * FROM computers WHERE name = ?", "M02B").Scan(&cn5)

	//COMPUTER_RESERVATION
	db.Model(&COMPUTER_RESERVATION{}).Create(&COMPUTER_RESERVATION{

		Date:     time.Now(),
		COMPUTER: cn1,
		TIME:     time_p1,
		User:     preecha,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

}
