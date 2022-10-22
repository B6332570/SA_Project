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
	database, err := gorm.Open(sqlite.Open("sa-05-gee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Employee{},
		&Role{},
		&Province{},
		&MemberClass{},
		&User{},
		&Computer_os{},
		&Computer_reservation{},
		&Computer{},
		&Time_com{},
	)

	db = database

	//User
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  "1111111111111",
		Address:   "ถนน a อำเภอ v",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	db.Model(&User{}).Create(&User{
		Pin:       "B6222222",
		FirstName: "kawin",
		LastName:  "l.pat",
		Civ:       "2222222222222",
		Phone:     "0922222222",
		Email:     "kawin@mail.com",
		Password:  "2222222222222",
		Address:   "หอ b อำเภอ r",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	//Computer_os data
	comp_os_name1 := Computer_os{
		Name: "Windows",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name1)

	comp_os_name2 := Computer_os{
		Name: "macOS",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name2)

	//Computer data
	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W01A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W02A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W03A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "M01B",
		Comp_room:   "ROOM B",
		Computer_os: comp_os_name2,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "M02B",
		Comp_room:   "ROOM B",
		Computer_os: comp_os_name2,
	})

	//TIME data
	time_p1 := Time_com{
		Time_com_period: "08:00 - 09:00",
	}
	db.Model(&Time_com{}).Create(&time_p1)

	time_p2 := Time_com{
		Time_com_period: "09:00 - 10:00",
	}
	db.Model(&Time_com{}).Create(&time_p2)

	time_p3 := Time_com{
		Time_com_period: "10:00 - 11:00",
	}
	db.Model(&Time_com{}).Create(&time_p3)

	time_p4 := Time_com{
		Time_com_period: "11:00 - 12:00",
	}
	db.Model(&Time_com{}).Create(&time_p4)

	time_p5 := Time_com{
		Time_com_period: "12:00 - 13:00",
	}
	db.Model(&Time_com{}).Create(&time_p5)

	time_p6 := Time_com{
		Time_com_period: "13:00 - 14:00",
	}
	db.Model(&Time_com{}).Create(&time_p6)

	time_p7 := Time_com{
		Time_com_period: "14:00 - 15:00",
	}
	db.Model(&Time_com{}).Create(&time_p7)

	time_p8 := Time_com{
		Time_com_period: "15:00 - 16:00",
	}
	db.Model(&Time_com{}).Create(&time_p8)

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ COMPUTER มาเก็บไว้ในตัวแปรก่อน
	// cn = comp_name ที่มาจาก COMP_NAME ใน Entity COMPUTER
	var cn1 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W01A").Scan(&cn1)
	var cn2 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W02A").Scan(&cn2)
	var cn3 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "W03A").Scan(&cn3)
	var cn4 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M01B").Scan(&cn4)
	var cn5 Computer
	db.Raw("SELECT * FROM computers WHERE COMP_NAME = ?", "M02B").Scan(&cn5)

	//COMPUTER_RESERVATION
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn1,
		Time_com: time_p1,
		User:     preecha,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

}
