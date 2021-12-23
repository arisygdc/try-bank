package util

import (
	"math/rand"
)

func RandString(leng int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	str := make([]rune, leng)
	for i := range str {
		str[i] = letter[rand.Intn(len(letter))]
	}
	return string(str)
}
