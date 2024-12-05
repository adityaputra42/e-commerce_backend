package helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const aplhabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(aplhabet)
	for i := 0; i < n; i++ {
		c := aplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomOwner() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", RandomString(6))
}
