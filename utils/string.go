package utils

import "math/rand"

var LettersSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var NumSet = []rune("0123456789")

var NumAndLettersSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(set []rune, length uint) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = set[rand.Intn(len(set))]
	}

	return string(b)
}
