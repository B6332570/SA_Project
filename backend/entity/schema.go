package entity

import (
	"time"

	"gorm.io/gorm"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

type Role struct {
	gorm.Model
	Name       string
	BorrowDay  int
	BookRoomHR int
	BookComHR  int
	Users      []User `gorm:"foreignKey:RoleID"`
}

type Province struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	Name     string
	Discount int
	Users    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
	gorm.Model
	Pin       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Civ       string `gorm:"uniqueIndex"`
	Phone     string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Address   string
	//FK
	EmployeeID    *uint
	RoleID        *uint
	ProvinceID    *uint
	MemberClassID *uint
	//JOIN
	Province    Province    `gorm:"references:id"`
	Role        Role        `gorm:"references:id"`
	MemberClass MemberClass `gorm:"references:id"`
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

type Computer_reservation struct {
	gorm.Model
	Date time.Time

	Computer_id *uint    //FK
	Computer    Computer `gorm:"references:id"` //JOIN

	// Computer_os_id *uint       //FK
	// Computer_os    Computer_os `gorm:"references:id"` //JOIN

	Time_com_id *uint    //FK
	Time_com    Time_com `gorm:"references:id"` //JOIN

	UserID *uint //FK
	User   User  `gorm:"references:id"` //JOIN
}

type Computer struct {
	gorm.Model
	Comp_name string
	Comp_room string

	Computer_os_id *uint       //FK
	Computer_os    Computer_os `gorm:"references:id"` //JOIN

	// COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:COMPUTER_ID"`
}

type Computer_os struct {
	gorm.Model
	Name string
	// COMPUTER []COMPUTER `gorm:"foreignKey:COMPUTER_OS_ID"`
}

type Time_com struct {
	gorm.Model
	Time_com_period      string
	Computer_reservation []Computer_reservation `gorm:"foreignKey:Time_com_id"`
}
