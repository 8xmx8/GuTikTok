package main

import (
	"GuTikTok/mdb"
	"fmt"
)

func main() {
	db := mdb.DB
	rdb := mdb.Rdb
	fmt.Println(db, rdb)
}
