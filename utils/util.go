package util

import (
	"math/rand"
	"time"
)

func RandomName(n int) string {
	var letters = []byte("absdadsadwqjsakdadsjaldjlsajdlsa")
	rand.Seed(time.Now().Unix())
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
