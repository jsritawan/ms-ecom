package model

import "time"

const (
	UserTableName = "users"
)

type User struct {
	Id        uint   `gorm:"column:id;primaryKey"`
	Code      string `gorm:"column:code;size:255;not null;unique"`
	Email     string `gorm:"column:email;size:255;not null;unique"`
	Password  string `gorm:"column:password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (User) TableName() string {
	return UserTableName
}
