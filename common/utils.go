package common

import (
	"math/rand"
	"time"
)

func RandomString(length int, prefix string) string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune,length)
	rand.Seed(time.Now().UnixNano()+ int64(rand.Intn(100)))
	for i := range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return prefix+string(b)
	
}