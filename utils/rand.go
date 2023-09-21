package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GetId 获取ID
func GetId(a, b int64) int64 {
	// 使用纳秒时间戳，保证递增
	now := time.Now()
	return now.UnixNano()*a - b
}

// RandString 获取随机字符串
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
