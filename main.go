package main

import (
	"GuTikTok/config"
	"GuTikTok/logging"
	"GuTikTok/mdb"
	"GuTikTok/utils/checks"
	"fmt"
)

func main() {

	logger := logging.Logger
	db := mdb.DB
	rdb := mdb.Rdb
	prefix := config.Conf.Consul.ConsulAnonymityPrefix
	fmt.Println(logger, db, rdb)
	fmt.Println(prefix)
	s := "Abc123!@#"
	is := checks.ValidatePassword(s, 8, 32)
	fmt.Println(is)
}
