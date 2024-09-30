package model

import "time"

const (
	ProfileTableName = "profiles"
)

type Profile struct {
	Id        uint   `gorm:"column:id;primaryKey"`
	FirstName string `gorm:"column:first_name;size:255;not null"`
	LastName  string `gorm:"column:last_name;size:255;not null"`
	Mobile    string `gorm:"column:mobile;size:10;not null"`
	Sex       string `gorm:"column:sex;size:1"`
	Status    string `gorm:"column:status;size:1"`
	Image     string `gorm:"column:image;size:255"`
	UserId    uint   `gorm:"column:user_id;foreignKey;references:users"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (Profile) TableName() string {
	return ProfileTableName
}
