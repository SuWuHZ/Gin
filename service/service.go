package service

import (
	"Gin/db"
	"Gin/models"
)

func GetUserInfo(workId string) (int, interface{}) {
	var userInfo models.User

	if rows := db.GetDB().Model(models.User{}).Where("work_id = ?", workId).First(&userInfo).RowsAffected; rows != 1 {
		return ErrUserNotExist, nil
	}

	return ErrNoError, userInfo
}

func AddOneUser(user * models.User)(int, interface{}) {
	if err := db.GetDB().Create(user).Error; err!= nil {
		return ErrDBInsertFailed, nil
	}

	return ErrNoError, nil
}
