package hashing

import (
	"crypto/rand"
	"math/big"
)

func GenerateHash() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	res := make([]byte, 8)
	length := big.NewInt(int64(len(letters)))

	for i := 0; i < 8; i++ {
		num, _ := rand.Int(rand.Reader, length)
		res[i] = letters[num.Int64()]
	}

	return string(res)
}
