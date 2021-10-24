package models

type User struct {
	ID        int64  `gorm:"column:id;PRIMARY_KEY,AUTO_INCREMENT" json:"id"`
	WordID    string `gorm:"column:word_id;UNIQUE_INDEX" json:"work_id"`
	Name      string `gorm:"column:name;index:name" json:"name"`
	Age       int    `gorm:"column:age" json:"age"`
	Sex       string `gorm:"column:sex" json:"sex"`
	TelNumber string `gorm:"column:tel_number" json:"tel_number"`
}

func (u *User) TableName() string {
	return "users"
}
