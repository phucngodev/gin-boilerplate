package model

import "time"

type User struct {
	BaseModel
	Name     string    `gorm:"column:name" json:"name"`
	Password string    `gorm:"column:password" json:"-"`
	Mobile   string    `gorm:"column:mobile" json:"mobile"`
	Email    string    `gorm:"column:email" json:"email"`
	Sex      uint      `gorm:"column:sex" json:"sex"`
	Age      uint      `gorm:"column:age" json:"age"`
	Birthday time.Time `gorm:"column:birthday" json:"birthday"`
}

func (User) TableName() string {
	return "user"
}
