package models

import (
	db "MeTube/database"
	"time"
)

type MUserInfo struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Username   string `gorm:"column:username" json:"username"`
	UserPasswd string `gorm:"column:userpasswd" json:"userpasswd"`
}

func (p *MUserInfo) TableName() string {
	return "userInfo"
}

func GetUser(uname string, passwd string) *MUserInfo {
	user := &MUserInfo{}

	err := db.ORM.Where("username = ? and userPasswd = ?", uname, passwd).First(user).Error
	if err != nil || user.Id <= 0 {
		return nil
	}
	return user
}