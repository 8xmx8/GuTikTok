package main

import (
	"GuTikTok/utils/logging"

	"GuTikTok/mdb"
	"fmt"
)

func main() {

	logger := logging.Logger
	db := mdb.DB
	rdb := mdb.Rdb
	fmt.Println(logger, db, rdb)
}
