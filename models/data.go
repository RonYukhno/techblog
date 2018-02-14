package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Data struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

type User struct {
	Id       int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Login    string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "roman:36317@/techblog")
	if err != nil {
		log.Panic(err)
	}
	// defer db.Close()
	log.Println("Connection Established")
	db.AutoMigrate(&User{})

}

func (data *Data) CheckPaswrd() bool {
	user_db := &User{}

	if err := db.Debug().Table("users").Where("login = ?", data.Login).First(&user_db).Error; err != nil {
		log.Println("ERROR DB:", err)
		return false
	} else {
		if user_db.Password == data.Password {
			return true
		} else {
			return false
		}
	}
}
