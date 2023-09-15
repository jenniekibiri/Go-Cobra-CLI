package pkg

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	emailCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	emailLength  = 10
)

// RandomEmail generates a random email address.
func RandomEmail() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	email := make([]byte, emailLength)
	for i := range email {
		email[i] = emailCharset[random.Intn(len(emailCharset))]
	}
	return fmt.Sprintf("%s@example.com", strings.ToLower(string(email)))
}

// RandomName generates a random name.
func RandomName() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	name := make([]byte, emailLength)
	for i := range name {
		name[i] = emailCharset[random.Intn(len(emailCharset))]
	}
	return strings.Title(string(name))
}
