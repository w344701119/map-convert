package com

import (
	"math/rand"
	"time"
)

//随机生成字符串
func RandLetter(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(25) + 65
		result[i] = byte(b)
	}
	return string(result)
}

//随机字符串
func RandDigital(length int) string {
	//实例化一个切片长度和预生成长度一样
	result := make([]byte, length, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = byte(r.Intn(9) + 48)
	}
	return string(result)
}

//随机字符串
func RandomString(length int) string {
	bytes := []byte{
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
		97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
		65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	}
	result := make([]byte, length, length)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[rnd.Intn(len(bytes)-1)]
	}
	return string(result)
}
