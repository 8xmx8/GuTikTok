package main

import (
	"GuTikTok/config"
	"fmt"
)

func main() {
	address := config.Conf.Consul.Address
	fmt.Println(address)
}
