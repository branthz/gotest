package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Num int64

type User struct {
	Id                int
	Age               int
	UserNum           Num
	Name              string        `sql:"size:128;not null"`
	Birthday          time.Time     // Time
	CreatedAt         time.Time     // CreatedAt: Time of record is created, will be insert automatically
	UpdatedAt         time.Time     // UpdatedAt: Time of record is updated, will be updated automatically
	Emails            []Email       // Embedded structs
	BillingAddress    Address       // Embedded struct
	BillingAddressID  sql.NullInt64 // Embedded struct's foreign key
	ShippingAddress   Address       // Embedded struct
	ShippingAddressId int64         // Embedded struct's foreign key
	CreditCard        CreditCard
	Latitude          float64
	Languages         []Language `gorm:"many2many:user_languages;"`
	CompanyID         *int
	Company           Company
	Role
	PasswordHash      []byte
	IgnoreMe          int64                 `sql:"-"`
	IgnoreStringSlice []string              `sql:"-"`
	Ignored           struct{ Name string } `sql:"-"`
	IgnoredPointer    *User                 `sql:"-"`
}
type CreditCard struct {
	ID        int8
	Number    string
	UserId    sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Email struct {
	Id        int16
	UserId    int
	Email     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Company struct {
	Id    int64
	Name  string
	Owner *User `sql:"-"`
}
type Role struct {
	Namer string
}

type Address struct {
	ID        int
	Address1  string
	Address2  string
	Post      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type Language struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_languages;"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/zl?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	//create table

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
		log.Fatalln(err)
		return
	}

	//insert
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	ret := db.NewRecord(user)
	log.Println(ret)

}
