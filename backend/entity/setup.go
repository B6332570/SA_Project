package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

	password3, _ := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password4, _ := bcrypt.GenerateFromPassword([]byte("1"), 14)

	//Role

	employee := Role{
		Name:       "Employee",
		BorrowDay:  5,
		BookRoomHR: 6,
		BookComHR:  6,
	}
	db.Model(&Role{}).Create(&employee)

	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password3),
		Address:   "ถนน a อำเภอ v",
		//FK

		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "B6222222",
		FirstName: "kawin",
		LastName:  "anpa",
		Civ:       "22222222222222",
		Phone:     "0811111111",
		Email:     "kawin@mail.com",
		Password:  string(password4),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	//Computer_os data
	comp_os_name1 := Computer_os{
		Name: "Windows",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name1)

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
		Comp_name:   "W04A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W05A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
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
	// cn = comp_name ที่มาจาก Comp_name ใน Entity Computer
	var cn1 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W01A").Scan(&cn1)
	var cn2 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W02A").Scan(&cn2)
	var cn3 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W03A").Scan(&cn3)
	var cn4 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W04A").Scan(&cn4)
	var cn5 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W05A").Scan(&cn5)

	//Computer_reservation
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn1,
		Time_com: time_p1,
		User:     preecha,
	})

	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn2,
		Time_com: time_p2,
		User:     kawin,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

}
