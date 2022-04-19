package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func createPassword() string {

	const chars = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*+_-="
	password := ""

	for i := 0; i < 9; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}

	return password

}
