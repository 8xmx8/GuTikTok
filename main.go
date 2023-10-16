package main

import (
	"GuTikTok/src/models"
	"GuTikTok/src/storage/database"
	"GuTikTok/utils/logging"
	"gopkg.in/hlandau/passlib.v1"
)

func main() {
	user := models.User{
		Name: "yjy123",
		Pawd: "123456789434yjy",
	}
	pwd, _ := passlib.Hash(user.Pawd)
	user.Pawd = pwd
	tx := database.Client.Create(&user)
	if tx.RowsAffected == 1 {
		logging.Logger.Info("创建用户成功")
	}
}
