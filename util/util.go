package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// strTime layout: yyyy-mm-dd
func StrToTime(strTime string) (time.Time, error) {
	return time.Parse("2006-1-2", strings.Trim(strTime, " "))
}

func RandString(leng int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	str := make([]rune, leng)
	for i := range str {
		str[i] = letter[rand.Intn(len(letter))]
	}
	return string(str)
}

func RandNum(leng int) int {
	var numb = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var num string
	for i := 0; i < leng; i++ {
		num = num + numb[rand.Intn(10)]
	}
	res, _ := strconv.Atoi(num)
	return res
}
