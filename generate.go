package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func createPassword(length int) string {

	const chars = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*+_-="
	password := ""

	for i := 0; i < length; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}

	return password

}
