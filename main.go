package main

import (
	"GuTikTok/config"
	"GuTikTok/mdb"
)

func main() {
	config.InitConf()
	mdb.InitLog()
	mdb.InitDb()
	mdb.InitRdb()
}
