package pkg

import (
	"math/rand"
	"time"
)

func CodGenerator(length int) string {
	rand.NewSource(time.Now().UnixNano())

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters))
		code[i] = characters[randomIndex]
	}

	return string(code)
}
