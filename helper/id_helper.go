package helper

import (
	"math/rand"
)

const charset = "0123456789"

func Generate(identifier string) string {
	var length int = 10

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return identifier + string(b)
}
