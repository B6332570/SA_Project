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
	PROVINCE    PROVINCE
	ROLE        ROLE
	MemberClass MemberClass
	EMP         EMPLOYEE
	CR          []COMPUTER_RESERVATION `gorm:"foreignKey:UserID"`
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

type COMPUTER_RESERVATION struct {
	gorm.Model
	Date time.Time

	COMP_ID *uint    //FK
	C       COMPUTER //JOIN

	COMP_OS_ID *uint       //FK
	CO         COMPUTER_OS //JOIN

	TIME_ID *uint //FK
	TIME    TIME  //JOIN

	UserID *uint //FK
	User   User  //JOIN
}

type COMPUTER struct {
	gorm.Model
	COMP_NAME string
	COMP_ROOM string

	COMP_OS_ID *uint       //FK
	CO         COMPUTER_OS //JOIN

	Place_Class_id *uint                  //FK
	PC             Place_Class            //JOIN
	CR             []COMPUTER_RESERVATION `gorm:"foreignKey:C_ID"`
}

type COMPUTER_OS struct {
	gorm.Model
	COMP_OS string
	C       []COMPUTER             `gorm:"foreignKey:COMPUTER_OSID"`
	CR      []COMPUTER_RESERVATION `gorm:"foreignKey:COMPUTER_OSID"`
}

type Place_Class struct {
	gorm.Model
	Place_Class_name string
	C                []COMPUTER `gorm:"foreignKey:Place_ClassID"`
}

type TIME struct {
	gorm.Model
	TIME_PERIOD string
	CR          []COMPUTER_RESERVATION `gorm:"foreignKey:TIMEID"`
}
