package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())

}

// RandomEmail generates a random email address.
func RandomEmail() string {
	return RandomString(10) + "@example.com"
}

func RandomString(n int) string {
	var stringBuilder strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}
	return stringBuilder.String()

}
