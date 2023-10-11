package main

import (
	"GuTikTok/config"
	"fmt"
	"gopkg.in/hlandau/passlib.v1"
)

func main() {
	address := config.Conf.Consul.Address
	fmt.Println(address)

	password, _ := hashPassword("123456789")

	fmt.Println(password)

	hash := checkPasswordHash("123456789", password)
	fmt.Println(hash)

}
func hashPassword(password string) (string, error) {

	pwd, err := passlib.Hash(password)
	return pwd, err
}

func checkPasswordHash(password, hash string) bool {

	newHash, err := passlib.Verify(password, hash)

	return err == nil && newHash == ""
}
