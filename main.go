package main

import (
	"GuTikTok/src/models"
	"GuTikTok/src/storage/database"
	"GuTikTok/utils/logging"
	"fmt"
)

func main() {

	logger := logging.Logger
	client := database.Client
	comment := models.Comment{}
	fmt.Println(comment, logger, client)
}
