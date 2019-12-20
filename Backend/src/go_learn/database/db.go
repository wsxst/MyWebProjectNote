package database

import "github.com/jinzhu/gorm"

var ORM *gorm.DB

func InitDB() {
	var err error
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("Connect failed...")
	}
	ORM = db
	// 打印sql详情
	db.LogMode(true)
}