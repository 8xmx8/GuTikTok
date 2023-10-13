package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	stringsLib "strings"
)

func main() {
	email := "2647369456@qq.com"
	e := fmt.Sprintf("https://cravatar.cn/avatar/%s?d=identicon", getEmailMD5(email))

	fmt.Println(e)
}

func getEmailMD5(email string) (md5String string) {
	// 将邮箱地址转换为小写形式
	lowerEmail := stringsLib.ToLower(email)
	// 创建 MD5 哈希对象
	hashed := md5.New()
	// 将小写的邮箱地址转换为字节数组，并进行哈希计算
	hashed.Write([]byte(lowerEmail))
	// 获取计算后的 MD5 哈希值的字节数组
	md5Bytes := hashed.Sum(nil)
	// 将字节数组转换为十六进制字符串表示形式
	md5String = hex.EncodeToString(md5Bytes)
	return
}
