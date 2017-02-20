package utils

import (
	"math/rand"
	"time"
)

//生成随机序列
func UUIDgen() string {
	uuid := []byte{}
	temp := "qwertyuiopasdfghjklzxcvbnm"
	temp1 := []byte(temp)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		uuid = append(uuid, temp1[rand.Intn(26)])
	}
	return string(uuid)
}