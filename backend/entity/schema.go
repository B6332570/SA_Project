package entity

import (
	"time"

	"gorm.io/gorm"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
type EMPLOYEE struct {
	gorm.Model
	NAME     string
	PASSWORD string
	USERS    []User `gorm:"foreignKey:EMP_ID"`
}

type ROLE struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []User `gorm:"foreignKey:ROLE_ID"`
}

type PROVINCE struct {
	gorm.Model
	NAME  string
	USERS []User `gorm:"foreignKey:PROVINCE_ID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT string
	USERS    []User `gorm:"foreignKey:MemberClass_ID"`
}

type User struct {
	gorm.Model
	PIN       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	CIV       string `gorm:"uniqueIndex"`
	PHONE     string
	EMAIL     string `gorm:"uniqueIndex"`
	PASSWORD  string
	ADDRESS   string
	//FK
	EMP_ID         *uint
	ROLE_ID        *uint
	PROVINCE_ID    *uint
	MemberClass_ID *uint
	//JOIN
	PROVINCE             PROVINCE
	ROLE                 ROLE
	MemberClass          MemberClass
	EMP                  EMPLOYEE
	COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:UserID"`
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

type COMPUTER_RESERVATION struct {
	gorm.Model
	Date time.Time

	COMP_ID  *uint    //FK
	COMPUTER COMPUTER `gorm:"references:id"` //JOIN

	COMP_OS_ID  *uint       //FK
	COMPUTER_OS COMPUTER_OS `gorm:"references:id"` //JOIN

	TIME_ID *uint //FK
	TIME    TIME  `gorm:"references:id"` //JOIN

	UserID *uint //FK
	User   User  `gorm:"references:id"` //JOIN
}

type COMPUTER struct {
	gorm.Model
	COMP_NAME string
	COMP_ROOM string

	COMP_OS_ID  *uint       //FK
	COMPUTER_OS COMPUTER_OS `gorm:"references:id"` //JOIN

	// Place_Class_id *uint                  //FK
	// PC             Place_Class            //JOIN
	COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:COMP_ID"`
}

type COMPUTER_OS struct {
	gorm.Model
	Name     string
	COMPUTER []COMPUTER `gorm:"foreignKey:COMP_OS_ID"`
}

// type Place_Class struct {
// 	gorm.Model
// 	NAME string
// 	C    []COMPUTER `gorm:"foreignKey:Place_ClassID"`
// }

type TIME struct {
	gorm.Model
	TIME_PERIOD          string
	COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:TIME_ID"`
}
