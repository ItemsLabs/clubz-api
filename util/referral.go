package util

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateReferralCode() string {
	//return StringWithCharset(6, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return StringWithCharset(6, "0123456789")
}
