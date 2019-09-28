package key

import (
	"math/rand"
	"time"
)

// Generate takes a lenght and generates a random string of characters
func Generate(length int) string {
	char := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "!@#$%^&*()" + "1234567890"
	b := make([]byte, length)

	for i := range b {
		b[i] = char[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(char))]
	}

	return string(b)
}
