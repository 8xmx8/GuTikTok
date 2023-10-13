package main

import (
	"GuTikTok/src/models"
	"GuTikTok/src/storage/database"
	"GuTikTok/src/storage/redis"
	"GuTikTok/utils/logging"
	"fmt"
)

func main() {

	logger := logging.Logger
	client := database.Client
	comment := models.Comment{}
	RedisCli := redis.Client
	fmt.Println(comment, logger, client, RedisCli)
}
