package model

import "time"

type BaseModel struct {
	Id            int64      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	EnabledStatus *int8      `gorm:"column:enabled_status;type:tinyint;default:1" json:"-"`
	Created       time.Time  `gorm:"column:created;type:datetime;default:CURRENT_TIMESTAMP" json:"created"`
	Modified      *time.Time `gorm:"column:modified;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified"`
}
