package util

import (
	crptRand "crypto/rand"
	"math/big"
	"math/rand"
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

func RandNum(max int64) int64 {
	n, _ := crptRand.Int(crptRand.Reader, big.NewInt(max))
	return n.Int64()
}
