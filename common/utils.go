package common

import "math/rand"

func RandomString(length int, prefix string) string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune,length)
	for i := range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return prefix+string(b)
	
}