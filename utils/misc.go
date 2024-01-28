package utils

import "math/rand"

func GetRandomCode(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomCode := ""
	for i := 0; i < length; i++ {
		randomCode += string(chars[rand.Intn(len(chars))])
	}
	return randomCode
}
