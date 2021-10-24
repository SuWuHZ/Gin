package service

import (
	"Gin/db"
	"Gin/models"
)

func GetUserInfo(userId string) (int, interface{}) {
	var userInfo models.User

	if err := db.GetDB().Model(models.User{}).Where("work_id = ?").First(&userInfo).Error; err != nil {
		return ErrDBQueryFailed, nil
	}

	return ErrNoError, userInfo
}
