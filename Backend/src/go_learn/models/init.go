package models

import (
	"go_learn/database"
)

func AutoMigrate() {
	database.ORM.AutoMigrate(&MUserInfo{})
}