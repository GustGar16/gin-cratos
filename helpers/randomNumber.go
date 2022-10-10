package helpers

import (
	"crypto/rand"
	"io"
)

func NewRandomNumber(limit int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	result := make([]byte, limit)
	n, err := io.ReadAtLeast(rand.Reader, result, limit)
	if n != limit {
		panic(err)
	}
	for i := 0; i < len(result); i++ {
		result[i] = table[int(result[i])%len(table)]
	}
	return string(result)
}
