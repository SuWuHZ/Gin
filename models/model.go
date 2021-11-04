package models

import "time"

type User struct {
	ID        int64     `gorm:"column:id;PRIMARY_KEY,AUTO_INCREMENT" json:"id"`
	WordID    string    `gorm:"column:work_id;UNIQUE_INDEX" json:"work_id"`
	Name      string    `gorm:"column:name;index:name" json:"name"`
	Age       int       `gorm:"column:age" json:"age"`
	Sex       string    `gorm:"column:sex" json:"sex"`
	TelNumber string    `gorm:"column:tel_number" json:"tel_number"`
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
}

func (User) TableName() string {
	return "users"
}
