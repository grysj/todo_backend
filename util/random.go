package util

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(chars)

	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUser() string {
	n := RandomInt(3, 8)
	return RandomString(int(n))
}

func RandomTitle() string {
	n := RandomInt(5, 20)
	return RandomString(int(n))
}
