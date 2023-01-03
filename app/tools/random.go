package tools

import (
	"math/rand"
)

func RandomURL(size int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	str := make([]rune, size)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
